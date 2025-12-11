package domain

import "time"

var TargetDateLayout = "2006-01-02T15:04:05.999Z"

type FilterPeriod struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type FilterState struct {
	Status          string `json:"status,omitempty"`
	StatusExt       string `json:"statusExt,omitempty"`
	IsStatusExtNull bool   `json:"isStatusExtNull,omitempty"`
}

type FilterTarget struct {
	EmissionDatePeriod         FilterPeriod  `json:"emissionDatePeriod,omitempty"`
	ApplicationDatePeriod      FilterPeriod  `json:"applicationDatePeriod,omitempty"`
	ProductionDatePeriod       FilterPeriod  `json:"productionDatePeriod,omitempty"`
	IntroducedDatePeriod       FilterPeriod  `json:"introducedDatePeriod,omitempty"`
	Gtins                      []string      `json:"gtins,omitempty"`
	ProducerInns               []string      `json:"producerInns,omitempty"`
	GeneralPackageTypes        []string      `json:"generalPackageTypes,omitempty"`
	TurnoverTypes              []string      `json:"turnoverTypes,omitempty"`
	States                     []FilterState `json:"states,omitempty"`
	TnVed                      string        `json:"tnVed,omitempty"`
	TnVed10                    string        `json:"tnVed10,omitempty"`
	EmissionTypes              []string      `json:"emissionTypes,omitempty"`
	IsAggregated               bool          `json:"isAggregated,omitempty"`
	ProductGroups              []string      `json:"productGroups,omitempty"`
	EliminationReasons         []string      `json:"eliminationReasons,omitempty"`
	PrVetDoc                   string        `json:"prVetDoc,omitempty"`
	HaveChildren               bool          `json:"haveChildren,omitempty"`
	ServiceProviderIDPresented bool          `json:"serviceProviderIdPresented,omitempty"`
	ServiceProviderTypes       []string      `json:"serviceProviderTypes,omitempty"`
	OrderIds                   []string      `json:"orderIds,omitempty"`
}

type FilterPagination struct {
	PerPage          int    `json:"perPage"`
	LastEmissionDate string `json:"lastEmissionDate"`
	Sgtin            string `json:"sgtin"`
	Direction        int    `json:"direction"`
}
type TargetFilter struct {
	Filter     FilterTarget     `json:"filter"`
	Pagination FilterPagination `json:"pagination"`
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
	EmissionDate       time.Time     `json:"emissionDate"`
	GeneralPackageType string        `json:"generalPackageType"`
	OwnerInn           string        `json:"ownerInn"`
	EmissionType       string        `json:"emissionType"`
	ProductGroup       string        `json:"productGroup"`
	HaveChildren       bool          `json:"haveChildren"`
	EliminationReason  string        `json:"eliminationReason"`
	ReceiptDate        time.Time     `json:"receiptDate"`
	Expiration         []interface{} `json:"expiration"`
	OrderID            string        `json:"orderId"`
	ModInfo            ModInfo       `json:"modInfo"`
}

type TargetResult struct {
	Result     []Result `json:"result"`
	IsLastPage bool     `json:"isLastPage"`
}

type TargetCis struct {
	Result
}
