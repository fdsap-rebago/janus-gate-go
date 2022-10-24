package response

type SuperSaving struct {
	// Client ID
	Cid int `json:"cid"`
	// Client Fullname
	FullName string `json:"fullname"`
	// Client Account
	Acc string `json:"acc"`
	// Account Balance
	Balance float32 `json:"balance"`
	// Withdrawable
	Withdrawable float32 `json:"withdrawable"`
	// Center Code
	CenterCode int `json:"centerCode"`
	// Unit Code
	UnitCode int `json:"unitCode"`
	// Center Name
	CenterName string `json:"centerName"`
	// Unit Name
	UnitName string `json:"unitName"`
	// Withdrawal Amount
	WithdrawalAmount float32 `json:"withdrawalAmount"`
}
