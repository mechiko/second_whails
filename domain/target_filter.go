package domain

type FilterTarget struct {
	// EmissionDatePeriod struct {
	// 	From string `json:"from"`
	// 	To   string `json:"to"`
	// } `json:"emissionDatePeriod,omitempty"`
	// ApplicationDatePeriod struct {
	// 	From string `json:"from"`
	// 	To   string `json:"to"`
	// } `json:"applicationDatePeriod,omitempty"`
	// ProductionDatePeriod struct {
	// 	From string `json:"from"`
	// 	To   string `json:"to"`
	// } `json:"productionDatePeriod,omitempty"`
	// IntroducedDatePeriod struct {
	// 	From string `json:"from"`
	// 	To   string `json:"to"`
	// } `json:"introducedDatePeriod,omitempty"`
	Gtins               []string `json:"gtins,omitempty"`
	ProducerInns        []string `json:"producerInns,omitempty"`
	GeneralPackageTypes []string `json:"generalPackageTypes,omitempty"`
	TurnoverTypes       []string `json:"turnoverTypes,omitempty"`
	// States              []struct {
	// 	Status          string `json:"status,omitempty"`
	// 	StatusExt       string `json:"statusExt,omitempty"`
	// 	IsStatusExtNull bool   `json:"isStatusExtNull,omitempty"`
	// } `json:"states,omitempty"`
	TnVed                      string   `json:"tnVed,omitempty"`
	TnVed10                    string   `json:"tnVed10,omitempty"`
	EmissionTypes              []string `json:"emissionTypes,omitempty"`
	IsAggregated               bool     `json:"isAggregated,omitempty"`
	ProductGroups              []string `json:"productGroups,omitempty"`
	EliminationReasons         []string `json:"eliminationReasons,omitempty"`
	PrVetDoc                   string   `json:"prVetDoc,omitempty"`
	HaveChildren               bool     `json:"haveChildren,omitempty"`
	ServiceProviderIDPresented bool     `json:"serviceProviderIdPresented,omitempty"`
	ServiceProviderTypes       []string `json:"serviceProviderTypes,omitempty"`
	OrderIds                   []string `json:"orderIds,omitempty"`
}

type TargetFilter struct {
	Filter     FilterTarget `json:"filter"`
	Pagination struct {
		PerPage          int    `json:"perPage"`
		LastEmissionDate string `json:"lastEmissionDate"`
		Sgtin            string `json:"sgtin"`
		Direction        int    `json:"direction"`
	} `json:"pagination"`
}

type ModInfo struct {
	ModID   int    `json:"modId"`
	Kpp     string `json:"kpp"`
	Address string `json:"address"`
}

type Result struct {
	Sgtin              string        `json:"sgtin"`
	Cis                string        `json:"cis"`
	CisWithoutBrackets string        `json:"cisWithoutBrackets"`
	Gtin               string        `json:"gtin"`
	ProducerInn        string        `json:"producerInn"`
	Status             string        `json:"status"`
	EmissionDate       string        `json:"emissionDate"`
	GeneralPackageType string        `json:"generalPackageType"`
	OwnerInn           string        `json:"ownerInn"`
	EmissionType       string        `json:"emissionType"`
	ProductGroup       string        `json:"productGroup"`
	HaveChildren       bool          `json:"haveChildren"`
	EliminationReason  string        `json:"eliminationReason"`
	ReceiptDate        string        `json:"receiptDate"`
	Expiration         []interface{} `json:"expiration"`
	OrderID            string        `json:"orderId"`
	ModInfo            ModInfo       `json:"modInfo"`
}

type TargetResult struct {
	Result     []Result `json:"result"`
	IsLastPage bool     `json:"isLastPage"`
}
