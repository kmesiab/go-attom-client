package models

type Location struct {
	Accuracy  string  `json:"accuracy"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Distance  int     `json:"distance"`
	GeoID     string  `json:"geoId"`
	GeoIDV4   GeoIDV4 `json:"geoIdV4"`
}
