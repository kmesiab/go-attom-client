package models

type Rooms struct {
	BathFixtures int `json:"bathfixtures"`
	BathsFull    int `json:"bathsfull"`
	BathsTotal   int `json:"bathsTotal,bathstotal"`
	Beds         int `json:"beds"`
	RoomsTotal   int `json:"roomsTotal"`
}
