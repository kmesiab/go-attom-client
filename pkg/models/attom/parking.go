package models

type Parking struct {
	GarageType string `json:"garagetype"`
	PrkgSize   int    `json:"prkgSize"`
	PrkgType   string `json:"prkgType"`
}
