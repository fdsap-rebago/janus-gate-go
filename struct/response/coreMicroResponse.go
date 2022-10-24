package response

type SuperSaving struct {
	Cid              int     `json:"cid"`
	FullName         string  `json:"fullname"`
	Acc              string  `json:"acc"`
	Balance          float32 `json:"balance"`
	Withdrawable     float32 `json:"withdrawable"`
	CenterCode       int     `json:"centerCode"`
	UnitCode         int     `json:"unitCode"`
	CenterName       string  `json:"centerName"`
	UnitName         string  `json:"unitName"`
	WithdrawalAmount float32 `json:"withdrawalAmount"`
}

type ReferenceCallBack struct {
	RetCode int    `json:"retCode"`
	Message string `json:"message"`
}

type CustSavingInfo struct {
	Cid                int         `json:"cid"`
	Acc                interface{} `json:"acc"`
	AppType            int         `json:"appType"`
	AcctType           int         `json:"acctType"`
	AccDesc            string      `json:"accDesc"`
	Dopen              string      `json:"dopen"`
	Status             int         `json:"status"`
	ClassificationCode int         `json:"classificationCode"`
	ClassificationType int         `json:"classificationType"`
	Balance            float64     `json:"balance"`
}

type CustSavingsLists struct {
	Cid        int     `json:"cid"`
	Acc        string  `json:"acc"`
	Acctype    int     `json:"acctType"`
	AccDesc    string  `json:"accDesc"`
	Dopen      string  `json:"dopen"`
	StatusDesc string  `json:"statusDesc"`
	Balance    float32 `json:"balance"`
	Status     int     `json:"status"`
}
