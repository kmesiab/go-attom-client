package nationalflooddata

type Waterbody struct {
	AreaSqKm string  `json:"areasqkm"`
	DistKm   float64 `json:"distkm"`
	GnisID   string  `json:"gnis_id"`
	Name     string  `json:"name"`
	ObjectID string  `json:"objectid"`
	OgcFid   int     `json:"ogc_fid"`
	State    string  `json:"state"`
}
