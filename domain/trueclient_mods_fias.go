package domain

type ModFiasInfo struct {
	Result []struct {
		Address        string   `json:"address"`
		Kpp            string   `json:"kpp"`
		FiasID         string   `json:"fiasId,omitempty"`
		Inn            string   `json:"inn"`
		ProductGroups  []string `json:"productGroups"`
		IsBlockedEgais bool     `json:"isBlockedEgais,omitempty"`
	} `json:"result"`
	Total    int  `json:"total"`
	NextPage bool `json:"nextPage"`
}
