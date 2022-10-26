package response

type ResponseModel struct {
	RetCode string      `json:"retCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseModelwoData struct {
	RetCode int    `json:"retCode"`
	Message string `json:"message"`
}
