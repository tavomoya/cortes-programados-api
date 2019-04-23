package models

import "time"

type QueryOptions struct {
	Skip  *int `json:"skip,omitempty"`
	Limit *int `json:"limit,omitempty"`
}

type QueryDateRange struct {
	Equal            *time.Time `json:"$eq,omitempty" bson:"$eq,omitempty"`
	LesserThan       *time.Time `json:"$lt,omitempty" bson:"$lt,omitempty"`
	GreaterThan      *time.Time `json:"$gt,omitempty" bson:"$gt,omitempty"`
	LesserThanEqual  *time.Time `json:"$lte,omitempty" bson:"$lte,omitempty"`
	GreaterThanEqual *time.Time `json:"$gte,omitempty" bson:"$gte,omitempty"`
}
