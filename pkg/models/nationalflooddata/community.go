package nationalflooddata

type Community struct {
	Firm            string  `json:"firm"`
	RegemerSanction string  `json:"regemer_sanction"`
	Tribal          string  `json:"tribal"`
	Notes           *string `json:"notes"`
	CommName        string  `json:"comm_name"`
	CommPart        bool    `json:"comm_part"`
	Fhbm            string  `json:"fhbm"`
	Curreff         string  `json:"curreff"`
}
