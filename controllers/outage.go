package controllers

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/models"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type OutageController struct {
	db *lib.DBLib
}

func NewOutageController(db *lib.DBLib) *OutageController {
	return &OutageController{
		db: db,
	}
}

func (o *OutageController) GetAllOutages() ([]*models.Outage, error) {

	// var query interfac
	res, err := o.db.Find(nil, nil)
	if err != nil {
		fmt.Println(":|", err)
		return nil, err
	}

	outages := make([]*models.Outage, 0)
	err = lib.ParseInterfaceToStruct(res, &outages)
	if err != nil {
		return nil, fmt.Errorf("Error parsing response to a slice: %v", err)
	}

	return outages, nil
}

func (o *OutageController) FilterOutages(query *models.OutageFilter) ([]*models.Outage, error) {

	req := make(map[string]interface{}, 0)
	err := lib.StructToMap(query, &req)
	if err != nil {
		return nil, err
	}

	res, err := o.db.Find(req, nil)
	if err != nil {
		return nil, err
	}

	outages := make([]*models.Outage, 0)
	err = lib.ParseInterfaceToStruct(res, &outages)
	if err != nil {
		return nil, fmt.Errorf("Error parsing response to a slice: %v", err)
	}

	return outages, nil
}

func (o OutageController) CreateOutage(outage *models.Outage) (*bson.ObjectId, error) {

	newID := bson.NewObjectId()
	outage.ID = newID

	err := o.db.Insert(outage)
	if err != nil {
		return nil, err
	}

	return &newID, nil
}
