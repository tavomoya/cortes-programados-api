package models

import (
	"time"
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
	Province      string
	Sectors       []string
	Date          time.Time
	StartTime     string
	EndTime       string
	StartTimeInt  float64
	EndTimeInt    float64
	AffectedZones []string
	Company       string
	Circuit       string
}
