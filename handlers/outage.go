package handlers

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/lib/http_res"
	"net/http"
)

type OutageHandler struct {
	db *lib.DBLib
}

func NewOutageHandler(db *lib.DBLib) *OutageHandler {
	return &OutageHandler{
		db: db,
	}
}

func (o *OutageHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	outages, err := o.db.GetAllOutages()
	if err != nil {
		http_res.ErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	http_res.OKResponse(w, outages)
}
