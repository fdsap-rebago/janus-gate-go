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
