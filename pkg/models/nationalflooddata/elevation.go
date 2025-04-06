package nationalflooddata

type Elevation struct {
	PropertyElevation       float64              `json:"propertyelevation"`
	FloodBaseFloodElevation []BaseFloodElevation `json:"flood.basefloodelevation"`
	Coastline               []Coastline          `json:"coastline"`
	Waterbody               []Waterbody          `json:"waterbody"`
	StormSurge              StormSurge           `json:"stormsurge"`
}
