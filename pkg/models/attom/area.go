package models

type Area struct {
	BlockNum       string `json:"blockNum"`
	Loctype        string `json:"loctype"`
	CountrySecSubd string `json:"countrysecsubd"`
	CountyUse1     string `json:"countyuse1"`
	MunCode        string `json:"muncode"`
	MunName        string `json:"munname"`
	SrvyRange      string `json:"srvyRange"`
	SrvySection    string `json:"srvySection"`
	SrvyTownship   string `json:"srvyTownship"`
	SubdName       string `json:"subdname"`
	TaxCodeArea    string `json:"taxcodearea"`
}
