package models

type Size struct {
	BldgSize          int    `json:"bldgsize"`
	GrossSize         int    `json:"grosssize"`
	GrossSizeAdjusted int    `json:"grosssizeadjusted"`
	GroundFloorSize   int    `json:"groundfloorsize"`
	LivingSize        int    `json:"livingsize"`
	SizeInd           string `json:"sizeInd"`
	UniversalSize     int    `json:"universalsize,universalSize"`
}
