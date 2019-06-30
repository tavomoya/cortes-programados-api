package handlers

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/lib/http_res"
	"cortes-programados-api/models"
	"cortes-programados-api/scrapers/circuits"
	"net/http"
)

type CircuitHandler struct {
	DB             *lib.DBLib
	CircuitScraper *circuits.CircuitsScraper
}

func NewCircuitHandler(db *models.DatabaseConfig) *CircuitHandler {

	db.Collection = "circuits"

	dbLib := lib.NewDBLib(db)
	handler := &CircuitHandler{
		DB:             dbLib,
		CircuitScraper: circuits.New(dbLib),
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
