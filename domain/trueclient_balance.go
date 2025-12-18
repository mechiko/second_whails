package domain

type Balance []struct {
	Balance        int `json:"balance,omitempty"`
	ContractID     int `json:"contractId,omitempty"`
	OrganisationID int `json:"organisationId"`
	ProductGroupID int `json:"productGroupId"`
}

