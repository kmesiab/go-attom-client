package nationalflooddata

type Result struct {
	Loma          []Loma             `json:"loma"`
	FloodFirmPan  []FloodFirmPan     `json:"flood.s_firm_pan"`
	FloodFldHazAr []FloodFieldHazard `json:"flood.s_fld_haz_ar"`
	FloodPolAr    []FloodPolAr       `json:"flood.s_pol_ar"`
	CensusBureau  CensusBureau       `json:"census_bureau"`
	Community     Community          `json:"community"`
	Elevation     Elevation          `json:"elevation"`
	Property      Property           `json:"property"`
}

type FloodData struct {
	ParcelAddress *string `json:"parceladdress"`
	Geocode       Geocode `json:"geocode"`
	Coords        Coords  `json:"coords"`
	Result        Result  `json:"result"`
	MatchType     *string `json:"match_type"`
}
