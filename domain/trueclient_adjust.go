package domain

type AdjustBody struct {
	ParticipantInn string       `json:"participantInn"`
	Codes          []AdjustStep `json:"codes"`
}

type PermitDoc struct {
	PermitDocNumber string `json:"permitDocNumber"`
	PermitDocDate   string `json:"permitDocDate"`
	PermitDocType   int    `json:"permitDocType"`
}

type PermitDocsAdjust struct {
	PermitDocsOperation int         `json:"permitDocsOperation"`
	PermitDocs          []PermitDoc `json:"permitDocs,omitempty"`
}

type AdjustStep struct {
	Code           []string `json:"code"`
	ProductionDate string   `json:"productionDate,omitempty"`
	ExpirationDate string   `json:"expirationDate,omitempty"`
	Tnved          string   `json:"tnved,omitempty"`
	*PermitDocsAdjust
}
