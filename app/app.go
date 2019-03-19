package app

import (
	"cortes-programados-api/handlers"
	"cortes-programados-api/models"
	"cortes-programados-api/scrapers/edeeste"
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

	// norte, err := edenorte.ReadOutageAnouncement()
	// if err != nil {
	// 	return err
	// }

	// sur, err := edesur.GetOutageAnouncement()
	// if err != nil {
	// 	return err
	// }

	// outages := append(norte, sur...)

	_, err := edeeste.ReadOutageAnouncement()
	if err != nil {
		fmt.Println("Errrr", err)
		return err
	}

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

	h := handlers.NewOutageHandler(dbConfig)

	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/outage", h.GetAll).Methods("GET")
	router.HandleFunc("/outage/filter", h.Filter).Methods("POST")

	listen := fmt.Sprintf(":%d", config.Port)

	return http.ListenAndServe(listen, gorillah.CombinedLoggingHandler(os.Stdout, router))
}
