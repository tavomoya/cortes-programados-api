package controllers

import (
	"cortes-programados-api/lib"
	"cortes-programados-api/models"
)

type CircuitController struct {
	db *lib.DBLib
}

func NewCircuitController(db *lib.DBLib) *CircuitController {
	return &CircuitController{
		db: db,
	}
}

func (c *CircuitController) GetCircuits(query *models.QueryCircuits) ([]*models.Circuit, error) {

	res, err := c.db.Find(query, nil)
	if err != nil {
		return nil, err
	}

	circuits := make([]*models.Circuit, 0)
	err = lib.ParseInterfaceToStruct(res, &circuits)
	if err != nil {
		return nil, err
	}

	return circuits, nil
}
