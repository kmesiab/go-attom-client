package nationalflooddata

type Property struct {
	Sqft                   string `json:"sqft"`
	YearBuilt              string `json:"yearbuilt"`
	PropertyUseDescription string `json:"propertyusedescription"`
	ConstructionDesc       string `json:"constructiondesc"`
	StoriesCount           string `json:"storiescount"`
	FireResistance         string `json:"fireresistance"`
	ParkingGarageType      string `json:"parkinggaragetype"`
	ParkingGarageArea      string `json:"parkinggaragearea"`
}
