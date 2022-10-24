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

type GenerateCol struct {
	AppType        int         `json:"appType"`
	Code           int         `json:"code"`
	Status         int         `json:"status"`
	StatusDesc     string      `json:"statusDesc"`
	Acc            string      `json:"acc"`
	Iiid           int         `json:"iiid"`
	Um             string      `json:"um"`
	ClientName     string      `json:"clientName"`
	CenterCode     int         `json:"centerCode"`
	CenterName     string      `json:"centerName"`
	ManCode        int         `json:"manCode"`
	Unit           string      `json:"unit"`
	AreaCode       int         `json:"areaCode"`
	Area           string      `json:"area"`
	StaffName      string      `json:"staffName"`
	Accttype       int         `json:"accttype"`
	AcctDesc       string      `json:"acctDesc"`
	Disbdate       string      `json:"disbdate"`
	DateStart      string      `json:"dateStart"`
	Maturity       string      `json:"maturity"`
	Principal      float64     `json:"principal"`
	Interest       float64     `json:"interest"`
	Gives          int         `json:"gives"`
	IbalPrin       float64     `json:"ibalPrin"`
	IbalInt        float64     `json:"ibalInt"`
	BalPrin        float64     `json:"balPrin"`
	BalInt         float64     `json:"balInt"`
	Amort          float64     `json:"amort"`
	DuePrin        float64     `json:"duePrin"`
	DueInt         float64     `json:"dueInt"`
	LoanBal        float64     `json:"loanBal"`
	SaveBal        float64     `json:"saveBal"`
	WaiveInt       float64     `json:"waiveInt"`
	UnpaidCtr      int         `json:"unpaidCtr"`
	Writtenoff     int         `json:"writtenoff"`
	Classification int         `json:"classification"`
	ClassDesc      string      `json:"classDesc"`
	Writeoff       int         `json:"writeoff"`
	Pay            float64     `json:"pay"`
	Withdraw       float64     `json:"withdraw"`
	Type           int         `json:"type"`
	Uuid           interface{} `json:"uuid"`
}

type LoanInfo struct {
	Cid                 int         `json:"cid"`
	Acc                 string      `json:"acc"`
	AppType             int         `json:"appType"`
	AcctType            int         `json:"acctType"`
	AccDesc             string      `json:"accDesc"`
	Dopen               string      `json:"dopen"`
	Domaturity          string      `json:"domaturity"`
	Term                int         `json:"term"`
	Weekspaid           int         `json:"weekspaid"`
	Status              int         `json:"status"`
	Principal           float64     `json:"principal"`
	Interest            float64     `json:"interest"`
	Others              float64     `json:"others"`
	Discounted          float64     `json:"discounted"`
	Netproceed          float64     `json:"netproceed"`
	Balance             float64     `json:"balance"`
	Prin                float64     `json:"prin"`
	Intr                float64     `json:"intr"`
	Oth                 float64     `json:"oth"`
	Penalty             float64     `json:"penalty"`
	Waivedint           float64     `json:"waivedint"`
	Disbby              string      `json:"disbby"`
	Approvby            string      `json:"approvby"`
	Cycle               int         `json:"cycle"`
	Frequency           int         `json:"frequency"`
	Annumdiv            int         `json:"annumdiv"`
	Lngrpcode           int         `json:"lngrpcode"`
	Proff               int         `json:"proff"`
	Fundsource          int         `json:"fundsource"`
	Conintrate          float64     `json:"conintrate"`
	Amortcond           int         `json:"amortcond"`
	Amortcondvalue      float64     `json:"amortcondvalue"`
	Classification_code int         `json:"classification_code"`
	Classification_type int         `json:"classification_type"`
	Remarks             interface{} `json:"remarks"`
	Amort               interface{} `json:"amort"`
	IsLumpsum           int         `json:"isLumpsum"`
	LoanID              interface{} `json:"loanID"`
	AmortList           []AmortList `json:"amortList"`
	Charges             interface{} `json:"charges"`
}

type AmortList struct {
	Dnum        int     `json:"dnum"`
	Acc         string  `json:"acc"`
	DueDate     string  `json:"dueDate"`
	InstFlag    int     `json:"instFlag"`
	Prin        float64 `json:"prin"`
	Intr        float64 `json:"intr"`
	Oth         float64 `json:"oth"`
	Penalty     float64 `json:"penalty"`
	EndBal      float64 `json:"endBal"`
	EndInt      float64 `json:"endInt"`
	EndOth      float64 `json:"endOth"`
	InstPd      float64 `json:"instPd"`
	PenPd       float64 `json:"penPd"`
	CarVal      float64 `json:"carVal"`
	UpInt       float64 `json:"upInt"`
	ServFee     float64 `json:"servFee"`
	PledgeAmort float64 `json:"pledgeAmort"`
}

type Charges struct {
	Acc    interface{} `json:"acc"`
	Charge []Charges   `json:"charges"`
}

type OpenPaymentResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Details []Details `json:"details"`
}

type Details struct {
	SourceId            interface{} `json:"sourceId"`
	SourceAccountNumber interface{} `json:"sourceAccountNumber"`
	CustomerNumber      string      `json:"customerNumber"`
	CustomerName        string      `json:"customerName"`
	TransactionAmount   float32     `json:"transactionAmount"`
	ReferenceNumber     string      `json:"referenceNumber"`
	CoreReference       string      `json:"coreReference"`
	PaidDate            interface{} `json:"paidDate"`
}

type CustomerInformation struct {
	Cid               int         `json:"cid"`
	Lastname          string      `json:"lastname"`
	FirstName         string      `json:"firstname"`
	MiddleName        string      `json:"middlename"`
	MaidenFName       string      `json:"maidenFname"`
	MaidenLName       string      `json:"maidenLName"`
	MaidenMName       string      `json:"maidenMName"`
	DoBirth           string      `json:"doBirth"`
	BirthPlace        string      `json:"birthPlace"`
	Sex               int         `json:"sex"`
	CivilStatus       int         `json:"civilStatus"`
	Title             int         `json:"title"`
	Status            int         `json:"status"`
	Classification    int         `json:"classification"`
	SubClassification int         `json:"subClassification"`
	Business          int         `json:"business"`
	DEntry            string      `json:"doEntry"`
	DoRecognized      string      `json:"doRecognized"`
	DoResigned        string      `json:"doResigned"`
	BrCode            int         `json:"brCode"`
	UnitCode          int         `json:"unitCode"`
	CenterCode        int         `json:"centerCode"`
	Dosri             int         `json:"dosri"`
	Reffered          string      `json:"reffered"`
	Remarks           string      `json:"remarks"`
	AccountNumber     string      `json:"accountNumber"`
	SearchName        interface{} `json:"searchName"`

	Address interface{} `json:"address"`
	Contact interface{} `json:"contact"`

	MemberMaidenFName string      `json:"memberMaidenFName"`
	MemberMaidenLName string      `json:"memberMaidenLName"`
	MemberMaidenMName string      `json:"memberMaidenMName"`
	UnitName          interface{} `json:"unitName"`
	CenterName        interface{} `json:"centerName"`
}

type Address struct {
	AddressID int `json:"addressID"`
	Iiid      int `json:"iiid"`

	AddressTypeID  int    `json:"addressTypeID"`
	AddressType    string `json:"addressType"`
	ProvinceID     int    `json:"provinceID"`
	Province       string `json:"province"`
	TownCity       string `json:"townCity"`
	TowncityID     int    `json:"towncityID"`
	BarangayID     int    `json:"barangayID"`
	Barangay       int    `json:"barangay"`
	ZipCode        int    `json:"zipCode"`
	AddressDetails string `json:"addressDetails"`
}

type Contact struct {
	Iiid          int    `json:"iiid"`
	Series        int    `json:"series"`
	ContactTypeId int    `json:"contactTypeId"`
	ContactType   string `json:"contactType"`
	Contact       string `json:"contact"`
}

type LoanList struct {
	Acc         string  `json:"acc"`
	Status      string  `json:"status"`
	DataRelease string  `json:"dateRelease"`
	AcctType    string  `json:"acctType"`
	Principal   float32 `json:"principal"`
	Interest    float32 `json:"interest"`
	Oth         float32 `json:"oth"`
	Balance     float32 `json:"balance"`
	Term        int     `json:"term"`
	PaidTerm    int     `json:"paidTerm"`
}
