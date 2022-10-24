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
