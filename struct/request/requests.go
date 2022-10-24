package request

type CID struct {
	// Client ID
	Cid int `json:"cid" example:"10724585"`
}

type PRNumber struct {
	// PR Number
	PrNumber string `json:"prNumber" example:"AAP36467-1638243443748"`
}

type Acc struct {
	// Account
	Acc string `json:"acc" example:"1012-0000-00060776"`
}

type MultiplePaymentRequest struct {
	Payment         interface{} `json:"payment"`
	Trndate         string      `json:"trndate"`
	OrNumber        int         `json:"orNumber"`
	PrNumber        string      `json:"prNumber"`
	SourceId        int         `json:"sourceId"`
	Username        string      `json:"username"`
	Particulars     string      `json:"particulars"`
	RemitterCID     string      `json:"remitterCID"`
	TotalCollection int         `json:"totalCollection"`
}

type Payment struct {
	Um             string  `json:"um"`
	Acc            string  `json:"acc" example:"4062-0000-00672656"`
	Pay            int     `json:"pay" example:"50"`
	Area           string  `json:"area" example:"Laguna 2"`
	Code           int     `json:"code" example:"2"`
	Iiid           int     `json:"iiid" example:"9652333"`
	Type           string  `json:"type" example:"0"`
	Unit           string  `json:"unit" example:"Head Office"`
	Uuid           string  `json:"uuid"`
	Amort          int     `json:"amort"`
	Gives          int     `json:"gives"`
	BalInt         float32 `json:"balInt"`
	DueInt         int     `json:"dueInt"`
	Status         int     `json:"status"`
	AppType        int     `json:"appType"`
	BalPrin        float32 `json:"balPrin"`
	DuePrin        int     `json:"duePrin"`
	IbalInt        float32 `json:"ibalInt"`
	LoanBal        int     `json:"loanBal"`
	ManCode        int     `json:"manCode"`
	SaveBal        int     `json:"saveBal"`
	AcctDesc       string  `json:"acctDesc"`
	Accttype       int     `json:"accttype"`
	AreaCode       int     `json:"areaCode"`
	Disbdate       string  `json:"disbdate"`
	IbalPrin       float32 `json:"ibalPrin"`
	Interest       int     `json:"interest"`
	Maturity       string  `json:"maturity"`
	WaiveInt       int     `json:"waiveInt"`
	Withdraw       int     `json:"withdraw"`
	Writeoff       int     `json:"writeoff"`
	ClassDesc      string  `json:"classDesc"`
	DateStart      string  `json:"dateStart"`
	Principal      int     `json:"principal"`
	StaffName      string  `json:"staffName"`
	UnpaidCtr      int     `json:"unpaidCtr"`
	CenterCode     int     `json:"centerCode"`
	CenterName     string  `json:"centerName"`
	ClientName     string  `json:"clientName"`
	StatusDesc     string  `json:"statusDesc"`
	Writtenoff     int     `json:"writtenoff"`
	Classification int     `json:"classification"`
}

type OpenPaymentRequest struct {
	AccountNumber       string `json:"accountNumber" example:"1129000022901288"`
	Amount              int    `json:"amount" example:"20"`
	TrnReference        string `json:"trnReference" example:"CIP123-trndate01"`
	Particulars         string `json:"particulars" example:"Client Initiated Payment - Weekly Due Payment"`
	Username            string `json:"username" example:"konek2CARD"`
	SourceSaveAcc       string `json:"sourceSaveAcc" example:"1012-0000-37622903"`
	TransFee            int    `json:"transFee" example:"0"`
	TransFeeParticulars string `json:"transFeeParticulars"`
}
