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

	query := &models.OutageFilter{}
	// var query interfac
	res, err := o.db.Find(query, nil)
	if err != nil {
		fmt.Println(":|", err)
		return nil, err
	}

	outages := make([]*models.Outage, 0)
	err = lib.ParseInterfaceToStruct(res, &outages)
	if err != nil {
		return nil, err
	}

	return outages, nil
}

func (o *OutageController) FilterOutages(query *models.OutageFilter) ([]*models.Outage, error) {

	res, err := o.db.Find(query, nil)
	if err != nil {
		return nil, err
	}

	outages := make([]*models.Outage, 0)
	err = lib.ParseInterfaceToStruct(res, &outages)
	if err != nil {
		return nil, err
	}

	return outages, nil
}

func (o OutageController) CreateOutage(outage *models.Outage) (*bson.ObjectId, error) {

	newID := bson.NewObjectId()
	outage.ID = newID

	err := o.db.Insert(outage)
	if err != nil {
		return nil, lib.NewAppError(
			fmt.Sprintf("Error creating new outage: %v", err),
			"OutageController.CreateOutage",
		)
	}

	return &newID, nil
}
