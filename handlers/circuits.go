package handlers

import (
	"cortes-programados-api/controllers"
	"cortes-programados-api/lib"
	"cortes-programados-api/lib/http_res"
	"cortes-programados-api/models"
	"cortes-programados-api/scrapers/circuits"
	"encoding/json"
	"fmt"
	"net/http"
)

type CircuitHandler struct {
	DB             *lib.DBLib
	CircuitScraper *circuits.CircuitsScraper
	controller     *controllers.CircuitController
}

func NewCircuitHandler(db *models.DatabaseConfig) *CircuitHandler {

	db.Collection = "circuits"

	dbLib := lib.NewDBLib(db, "circuits")
	handler := &CircuitHandler{
		DB:             dbLib,
		CircuitScraper: circuits.New(dbLib),
		controller:     controllers.NewCircuitController(dbLib),
	}

	return handler
}

func (c *CircuitHandler) LoadCircuits(w http.ResponseWriter, r *http.Request) {

	err := c.CircuitScraper.ScrapeAllCircuits()
	if err != nil {
		http_res.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http_res.OKResponse(w, nil)
}

func (c *CircuitHandler) GetCircuits(w http.ResponseWriter, r *http.Request) {

	req := &models.QueryCircuits{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http_res.ErrorResponse(
			w,
			fmt.Errorf("Could not parse request body: %v", err),
			http.StatusBadRequest,
		)
		return
	}

	cs, err := c.controller.GetCircuits(req)
	if err != nil {
		http_res.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http_res.OKResponse(w, cs)
}
