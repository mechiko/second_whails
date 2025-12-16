package domain

type ApiGtinInfo struct {
	Results   []GtinInfo `json:"results"`
	ErrorCode string     `json:"errorCode"`
	Total     int        `json:"total"`
}

type GtinInfo struct {
	Name                     string `json:"name"`
	Gtin                     string `json:"gtin"`
	Brand                    string `json:"brand"`
	PackageType              string `json:"packageType"`
	InnerUnitCount           int    `json:"innerUnitCount"`
	Inn                      string `json:"inn"`
	ProductGroupID           int    `json:"productGroupId"`
	ProductGroup             string `json:"productGroup"`
	GoodSignedFlag           bool   `json:"goodSignedFlag"`
	GoodMarkFlag             bool   `json:"goodMarkFlag"`
	GoodTurnFlag             bool   `json:"goodTurnFlag"`
	IsKit                    bool   `json:"isKit"`
	IsTechGtin               bool   `json:"isTechGtin"`
	IsSet                    bool   `json:"isSet"`
	Level                    string `json:"level"`
	Multiplier               int    `json:"multiplier"`
	GoodStatus               string `json:"goodStatus"`
	FirstSignDate            int64  `json:"firstSignDate"`
	NcCreateDate             int64  `json:"ncCreateDate"`
	IsTonicDrink             string `json:"isTonicDrink"`
	CarbohydratesAmount      string `json:"carbohydratesAmount"`
	CoreVolume               int    `json:"coreVolume"`
	IsCarbonDioxideContains  string `json:"isCarbonDioxideContains"`
	IsSweetenerContains      string `json:"isSweetenerContains"`
	CountryAlpha2            string `json:"countryAlpha2"`
	Country                  string `json:"country"`
	TnVedCode                string `json:"tnVedCode"`
	TnVedCode10              string `json:"tnVedCode10"`
	Structure                string `json:"structure"`
	PackMaterial             string `json:"packMaterial"`
	BabyFoodProduct          string `json:"babyFoodProduct"`
	CarbonationMethod        string `json:"carbonationMethod"`
	ProductType              string `json:"productType"`
	Okpd2Group               string `json:"okpd2Group"`
	FullName                 string `json:"fullName"`
	IsSpecializedFoodProduct bool   `json:"isSpecializedFoodProduct"`
	VolumeProduct            string `json:"volumeProduct"`
	Taste                    string `json:"taste"`
	Fts                      struct {
		Rds []struct {
			AuthDocDate   string `json:"authDocDate"`
			AuthDocNumber string `json:"authDocNumber"`
		} `json:"rds"`
		Countries   []string `json:"countries"`
		TnVedCode10 []string `json:"tnVedCode10"`
	} `json:"fts"`
	PackageCharacteristic string `json:"packageCharacteristic"`
}
