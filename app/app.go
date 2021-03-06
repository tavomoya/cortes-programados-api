package app

import (
	"cortes-programados-api/handlers"
	"cortes-programados-api/lib/cron"
	"cortes-programados-api/models"
	"fmt"
	"net/http"
	"os"

	gorillah "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getDBSession(conn string) (*mgo.Session, error) {
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func Main(config *models.Config) error {

	session, err := getDBSession(config.ConnectionString)
	if err != nil {
		return err
	}

	defer session.Close()

	dbConfig := &models.DatabaseConfig{
		DB:               session.DB(config.DatabaseName),
		DatabaseName:     config.DatabaseName,
		ConnectionString: config.ConnectionString,
	}

	// CRON Job:
	job := cron.NewJob(dbConfig)
	job.UpdateOutagesCollection(config.Schedule)

	h := handlers.NewOutageHandler(dbConfig)
	circuitHandler := handlers.NewCircuitHandler(dbConfig)

	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/outage", h.GetAll).Methods("GET")
	router.HandleFunc("/outage/filter", h.Filter).Methods("POST")
	router.HandleFunc("/circuits/filter", circuitHandler.GetCircuits).Methods("POST")
	router.HandleFunc("/run/scrapers", h.RunScrapers).Methods("GET")
	router.HandleFunc("/load/circuits", circuitHandler.LoadCircuits).Methods("GET")

	listen := fmt.Sprintf(":%d", config.Port)

	return http.ListenAndServe(listen, gorillah.CombinedLoggingHandler(os.Stdout, router))
}
