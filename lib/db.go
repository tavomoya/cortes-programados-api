package lib

import (
	"cortes-programados-api/models"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type DBLib struct {
	config *models.DatabaseConfig
}

func NewDBLib(config *models.DatabaseConfig) *DBLib {
	return &DBLib{
		config: config,
	}
}

func (d *DBLib) FindbyID(id bson.ObjectId) (interface{}, error) {

	response := new(interface{})

	err := d.config.DB.C(d.config.Collection).FindId(id).One(response)
	if err != nil {
		return nil, fmt.Errorf("There was an error trying to get a record with that id: %v", err)
	}

	return response, nil
}

func (d *DBLib) Find(query *models.OutageFilter, options *models.QueryOptions) ([]interface{}, error) {

	var response []interface{}

	if options != nil {

		if options.Skip != nil && options.Limit != nil {

			err := d.config.DB.C(d.config.Collection).Find(query).Skip(*options.Skip).Limit(*options.Limit).All(&response)
			if err != nil {
				return nil, fmt.Errorf("There was an error trying to get a response with paginated query: %v", err)
			}

		}

	}

	err := d.config.DB.C(d.config.Collection).Find(query).All(&response)
	if err != nil {
		return nil, fmt.Errorf("There was an error trying to get a response with that query: %v", err)
	}

	return response, nil
}

func (d *DBLib) Insert(obj interface{}) error {

	err := d.config.DB.C(d.config.Collection).Insert(obj)
	if err != nil {
		return fmt.Errorf("There was an error trying to create a new record: %v", err)
	}

	return nil
}

func (d *DBLib) Update(id bson.ObjectId, obj interface{}) error {

	selector := bson.M{"_id": id}

	err := d.config.DB.C(d.config.Collection).Update(selector, obj)
	if err != nil {
		return fmt.Errorf("There was an error trying to update record %v: %v", id, err)
	}

	return nil
}

func (d *DBLib) Delete(id bson.ObjectId) error {

	selector := bson.M{"_id": id}

	err := d.config.DB.C(d.config.Collection).Remove(selector)
	if err != nil {
		return fmt.Errorf("There was an error trying to delete record %v: %v", id, err)
	}

	return nil
}
