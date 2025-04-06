package models

type Building struct {
	Size         Size            `json:"size"`
	Rooms        Rooms           `json:"rooms"`
	Interior     Interior        `json:"interior"`
	Construction Construction    `json:"construction"`
	Parking      Parking         `json:"parking"`
	Summary      BuildingSummary `json:"summary"`
}

type BuildingSize struct {
	BldgSize          int    `json:"bldgsize"`
	GrossSize         int    `json:"grosssize"`
	GrossSizeAdjusted int    `json:"grosssizeadjusted"`
	GroundFloorSize   int    `json:"groundfloorsize"`
	LivingSize        int    `json:"livingsize"`
	SizeInd           string `json:"sizeInd"`
	UniversalSize     int    `json:"universalsize"`
}

type BuildingSummary struct {
	ArchStyle          string `json:"archStyle"`
	Levels             int    `json:"levels"`
	Quality            string `json:"quality"`
	UnitsCount         string `json:"unitsCount"`
	View               string `json:"view"`
	YearBuiltEffective int    `json:"yearbuilteffective"`
}
