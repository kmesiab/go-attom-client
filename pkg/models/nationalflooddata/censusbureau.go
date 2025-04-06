package nationalflooddata

type CensusBureau struct {
	CensusBlock string `json:"census_block"`
	CBSA        CBSA   `json:"cbsa"`
	MetDiv      MetDiv `json:"metdiv"`
}
