package main

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/scrappers/edenorte"
	"cortes-programados-api/scrappers/edesur"
	"fmt"
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

	norte, err := edenorte.ReadOutageAnouncement()
	if err != nil {
		fmt.Println("Err")
		log.Fatal(err)
	}

	sur, err := edesur.GetOutageAnouncement()
	if err != nil {
		log.Fatal(err)
	}

	outages := append(norte, sur...)

	err = lib.InsertOuatageList(outages)
	if err != nil {
		log.Fatal(err)
	}

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
