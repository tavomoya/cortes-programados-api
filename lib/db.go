package lib

import (
	"cortes-programados-api/models"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type DBLib struct {
	config         *models.DatabaseConfig
	collectionName string
}

func NewDBLib(config *models.DatabaseConfig, collectionName string) *DBLib {
	return &DBLib{
		config:         config,
		collectionName: collectionName,
	}
}

func (d *DBLib) FindbyID(id bson.ObjectId) (interface{}, error) {

	response := new(interface{})

	err := d.config.DB.C(d.collectionName).FindId(id).One(response)
	if err != nil {
		return nil, fmt.Errorf("There was an error trying to get a record with that id: %v", err)
	}

	return response, nil
}

func (d *DBLib) Find(query interface{}, options *models.QueryOptions) ([]interface{}, error) {

	var response []interface{}

	if options != nil {

		if options.Skip != nil && options.Limit != nil {

			err := d.config.DB.C(d.collectionName).Find(query).Skip(*options.Skip).Limit(*options.Limit).All(&response)
			if err != nil {
				return nil, fmt.Errorf("There was an error trying to get a response with paginated query: %v", err)
			}

		}

	}

	err := d.config.DB.C(d.collectionName).Find(query).All(&response)
	if err != nil {
		return nil, fmt.Errorf("There was an error trying to get a response with that query: %v", err)
	}

	return response, nil
}

func (d *DBLib) Insert(obj interface{}) error {

	err := d.config.DB.C(d.collectionName).Insert(obj)
	if err != nil {
		return fmt.Errorf("There was an error trying to create a new record: %v", err)
	}

	return nil
}

func (d *DBLib) Update(id bson.ObjectId, obj interface{}) error {

	selector := bson.M{"_id": id}

	err := d.config.DB.C(d.collectionName).Update(selector, obj)
	if err != nil {
		return fmt.Errorf("There was an error trying to update record %v: %v", id, err)
	}

	return nil
}

func (d *DBLib) Delete(id bson.ObjectId) error {

	selector := bson.M{"_id": id}

	err := d.config.DB.C(d.collectionName).Remove(selector)
	if err != nil {
		return fmt.Errorf("There was an error trying to delete record %v: %v", id, err)
	}

	return nil
}
