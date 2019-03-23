package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Outage represents a scheduled electricity cut from the
// company of the region (EDENorte, EDESur, EDEEste)
// --
// Province - The province where the outage is going to take place
// Sector - A list of neighborhoods inside the province that will be affected
// Date -  The date in which the outage will take place
// Start - Start time of the outage
// End - End time of the outage
// AffectedZones - List of specific places within the Province/Sector that will be affected (This may include street names, and specific neighborhoods)
// Company - The company that scheduled the outage
// Circuit - The code of the electrical circuit of the affected zone. These are defined by the CDEEE
type Outage struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	Province      string        `json:"province"`
	Sectors       []string      `json:"sectors"`
	Date          time.Time     `json:"date"`
	StartTime     string        `json:"startTime"`
	EndTime       string        `json:"endTime"`
	AffectedZones []string      `json:"affectedZones"`
	Company       string        `json:"company"`
	Circuit       string        `json:"circuit"`
	StartTimeInt  float64
	EndTimeInt    float64
}

type OutageFilter struct {
	Company *string    `json:"company,omitempty"`
	Circuit *string    `json:"circuit,omitempty"`
	Date    *time.Time `json:"date,omitempty"`
}
