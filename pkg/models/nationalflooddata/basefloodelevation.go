package nationalflooddata

type BaseFloodElevation struct {
	BfeLnID   *string `json:"bfe_ln_id"`
	BfeType   string  `json:"bfe_type"`
	DfirmID   string  `json:"dfirm_id"`
	DistKm    int     `json:"distkm"`
	Elevation string  `json:"elevation"`
	FldArID   string  `json:"fld_ar_id"`
	FldZone   string  `json:"fld_zone"`
	LenUnit   string  `json:"len_unit"`
	VDatum    string  `json:"v_datum"`
	ZoneSubty *string `json:"zone_subty"`
}
