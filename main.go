package main

import (
	"cortes-programados-api/scrappers/edesur"
	"log"
	"net/http"
	"os"

	gorillah "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	// norte, err := edenorte.ReadOutageAnouncement()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = lib.InsertOuatageList(norte)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	edesur.GetOutageAnouncement()

	router := mux.NewRouter()

	router.HandleFunc("/", healthCheck)

	listen := os.Getenv("PORT")

	if listen == "" {
		listen = "9000"
	}

	if err := http.ListenAndServe(":"+listen, gorillah.CombinedLoggingHandler(os.Stdout, router)); err != nil {
		log.Fatal(err)
	}
}
