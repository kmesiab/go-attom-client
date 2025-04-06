package models

type Address struct {
	Country     string `json:"country"`
	CountrySubd string `json:"countrySubd"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Locality    string `json:"locality"`
	MatchCode   string `json:"matchCode"`
	OneLine     string `json:"oneLine"`
	Postal1     string `json:"postal1"`
	Postal2     string `json:"postal2"`
	Postal3     string `json:"postal3"`
}
