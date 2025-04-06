package nationalflooddata

type Geocode struct {
	Relevance  int     `json:"relevance"`
	MatchLevel string  `json:"matchLevel"`
	Label      string  `json:"label"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
