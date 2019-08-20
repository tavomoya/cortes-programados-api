package handlers

import (
	"cortes-programados-api/controllers"
	"cortes-programados-api/lib"
	"cortes-programados-api/lib/cron"
	"cortes-programados-api/lib/http_res"
	"cortes-programados-api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type OutageHandler struct {
	DB         *lib.DBLib
	dbConfig   *models.DatabaseConfig
	controller *controllers.OutageController
}

func NewOutageHandler(db *models.DatabaseConfig) *OutageHandler {

	db.Collection = "outages"

	dbLib := lib.NewDBLib(db, "outages")
	handler := &OutageHandler{
		DB:         dbLib,
		dbConfig:   db,
		controller: controllers.NewOutageController(dbLib),
	}

	return handler
}

func (o *OutageHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	outages, err := o.controller.GetAllOutages()
	if err != nil {
		http_res.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http_res.OKResponse(w, outages)
}

func (o *OutageHandler) Filter(w http.ResponseWriter, r *http.Request) {

	req := &models.OutageFilter{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http_res.ErrorResponse(
			w,
			fmt.Errorf("Could not parse request body: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	outages, err := o.controller.FilterOutages(req)
	if err != nil {
		http_res.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http_res.OKResponse(w, outages)
}

func (o *OutageHandler) RunScrapers(w http.ResponseWriter, r *http.Request) {

	outages, err := cron.GetOutagesScrapeData(o.dbConfig)
	if err != nil {
		http_res.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// db := lib.NewDBLib(dbConfig)
	err = cron.SaveOutageCollection(o.DB, outages)
	if err != nil {
		http_res.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http_res.OKResponse(w, nil)
}
