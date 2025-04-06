package models

type BuildingPermit struct {
	EffectiveDate string   `json:"effectiveDate"`
	PermitNumber  string   `json:"permitNumber"`
	Status        string   `json:"status,omitempty"`
	Description   string   `json:"description,omitempty"`
	Type          string   `json:"type,omitempty"`
	ProjectName   string   `json:"projectName,omitempty"`
	JobValue      int      `json:"jobValue,omitempty"`
	Fees          int      `json:"fees"`
	BusinessName  string   `json:"businessName"`
	HomeOwnerName string   `json:"homeOwnerName"`
	Classifiers   []string `json:"classifiers"`
}
