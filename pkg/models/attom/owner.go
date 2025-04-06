package models

type Owner struct {
	CorporateIndicator         string    `json:"corporateindicator"`
	Owner1                     OwnerInfo `json:"owner1"`
	Owner2                     OwnerInfo `json:"owner2"`
	Owner3                     OwnerInfo `json:"owner3"`
	Owner4                     OwnerInfo `json:"owner4"`
	OwnerRelationshipRightCode string    `json:"ownerrelationshiprightscode"`
	OwnerRelationshipType      string    `json:"ownerrelationshiptype"`
	AbsenteeOwnerStatus        string    `json:"absenteeownerstatus"`
	MailingAddressOneLine      string    `json:"mailingaddressoneline"`
}

type OwnerInfo struct {
	FullName       string `json:"fullname"`
	LastName       string `json:"lastname"`
	FirstNameAndMI string `json:"firstnameandmi"`
}

// OwnerDetail helps map OwnerInfo when it is sometimes named
// differently depending on the endpoint
type OwnerDetail OwnerInfo
