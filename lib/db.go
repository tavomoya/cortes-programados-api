package lib

import (
	"cortes-programados-api/models"
	"log"
	"os"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

func getDBSession() (*mgo.Session, error) {
	conn := os.Getenv("CONN_STRING")
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}

	return session, nil
}

// InsertOuatageList receives a list of outages to be inserted into the database
// --
// Params
// --
// outages - List of outages to be inserted
func InsertOuatageList(outages []*models.Outage) error {
	db, err := getDBSession()
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer db.Close()

	for _, o := range outages {
		err = db.DB("cortes-programados").C("outages").Insert(o)
		if err != nil {
			if !strings.Contains(err.Error(), "outage_unq dup key") {
				return err
			}
		}
	}

	return nil
}
