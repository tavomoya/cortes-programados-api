package models

// Circuit represents the information of an electrical circuit
// as given by the CDEEE (company in charge of regulating electricity in the country)
// --
// Name - Code name that represents the circuit. This is unique
// Company - The company to which the circuit belongs to
// Province - The province where the circuit is located
// CircuitType - Classifies circuits by how much outages/power they get per day
// InfluenceZones - List of places of influence inside the province
type Circuit struct {
	Name           string   `json:"name,omitempty" bson:"name,omitempty"`
	Company        string   `json:"company,omitempty" bson:"company,omitempty"`
	Province       string   `json:"province,omitempty" bson:"province,omitempty"`
	CircuitType    string   `json:"circuit_type,omitempty" bson:"circuit_type,omitempty"`
	InfluenceZones []string `json:"influence_zones,omitempty" bson:"influence_zones,omitempty"`
}
