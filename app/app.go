package app

import (
	"cortes-programados-api/handlers"
	"cortes-programados-api/lib"
	"cortes-programados-api/models"
	"cortes-programados-api/scrapers/edenorte"
	"cortes-programados-api/scrapers/edesur"
	"fmt"
	"net/http"
	"os"

	gorillah "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Main(config *models.Config) error {

	norte, err := edenorte.ReadOutageAnouncement()
	if err != nil {
		return err
	}

	sur, err := edesur.GetOutageAnouncement()
	if err != nil {
		return err
	}

	outages := append(norte, sur...)

	db, err := lib.NewDBLib(config)
	if err != nil {
		return err
	}

	defer db.Session.Close()

	err = db.InsertOuatageList(outages)
	if err != nil {
		return err
	}

	h := handlers.NewOutageHandler(db)

	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/outage", h.GetAll).Methods("GET")

	listen := fmt.Sprintf(":%d", config.Port)

	return http.ListenAndServe(listen, gorillah.CombinedLoggingHandler(os.Stdout, router))
}
