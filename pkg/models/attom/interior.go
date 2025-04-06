package models

type Interior struct {
	BsmtSize  int    `json:"bsmtsize"`
	FplcCount int    `json:"fplccount"`
	FplcInd   string `json:"fplcind"`
	FplcType  string `json:"fplctype"`
}
