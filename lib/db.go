package lib

import (
	"cortes-programados-api/models"
	"fmt"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

func getDBSession(conn string) (*mgo.Session, error) {
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}

	return session, nil
}

type DBLib struct {
	config  *models.Config
	Session *mgo.Session
}

func NewDBLib(config *models.Config) (*DBLib, error) {
	if config == nil || config.ConnectionString == "" {
		return nil, fmt.Errorf("No connection string supplied")
	}

	db := &DBLib{
		config: config,
	}

	session, err := getDBSession(config.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %v", err)
	}

	db.Session = session

	return db, nil
}

// InsertOuatageList receives a list of outages to be inserted into the database
// --
// Params
// --
// outages - List of outages to be inserted
func (d *DBLib) InsertOuatageList(outages []*models.Outage) error {

	for _, o := range outages {
		err := d.Session.DB("cortes-programados").C("outages").Insert(o)
		if err != nil {
			if !strings.Contains(err.Error(), "outage_unq dup key") {
				return err
			}
		}
	}

	return nil
}

func (d *DBLib) GetAllOutages() ([]*models.Outage, error) {

	outages := make([]*models.Outage, 0)
	err := d.Session.DB("cortes-programados").C("outages").Find(nil).All(&outages)
	if err != nil {
		return nil, fmt.Errorf("There was an error trying to get data: %v", err)
	}

	return outages, nil
}
