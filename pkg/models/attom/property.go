package models

type PropertyDetail struct {
	Status   Status     `json:"status"`
	Property []Property `json:"property"`
}

type Property struct {
	Identifier      Identifier       `json:"identifier"`
	Lot             Lot              `json:"lot"`
	Address         Address          `json:"address"`
	Location        Location         `json:"location"`
	Summary         PropertySummary  `json:"summary"`
	Building        Building         `json:"building"`
	BuildingPermits []BuildingPermit `json:"buildingPermits"`
	Vintage         Vintage          `json:"vintage"`
}

type Identifier struct {
	ID      int    `json:"Id"`
	Fips    string `json:"fips"`
	Apn     string `json:"apn"`
	AttomID int    `json:"attomId"`
}

type PropertySummary struct {
	PropClass     string `json:"propClass"`
	PropSubType   string `json:"propSubType"`
	PropType      string `json:"propType"`
	PropertyType  string `json:"propertyType"`
	YearBuilt     int    `json:"yearBuilt"`
	PropLandUse   string `json:"propLandUse"`
	PropIndicator int    `json:"propIndicator"`
}
