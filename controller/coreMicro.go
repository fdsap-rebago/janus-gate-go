package coremicro

import (
	"JanusGate/middleware"
	"JanusGate/struct/request"
	"JanusGate/struct/response"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "JanusGate/docs"

	"github.com/gofiber/fiber/v2"
)

// Developer		rebago
// @summary 	  	CORE-AVAILABLE-BALANCE
// @Description	  	Used to get the total and available amount of client savings
// @Tags		  	CORE MICRO
// @Accept		  	json
// @Produce		  	json
// @Param       	cid body request.CID true  "Client ID"
// @Success		  	200 {object} response.SuperSaving
// @Failure		  	400 {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/getSavingsForSuperApp [post]
func SuperAppSaving(c *fiber.Ctx) error {
	rcid := request.CID{}

	if parsErr := c.BodyParser(&rcid); parsErr != nil {
		middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", rcid, "CORE-MICRO", nil, parsErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	marshallReq, marshallErr := json.Marshal(rcid.Cid)

	fmt.Println("Marshall:", marshallErr)
	if marshallErr != nil {
		middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", rcid, "CORE-MICRO", nil, marshallErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", "application/json; charset=utf-8", bytes.NewBuffer(marshallReq))

	if respErr != nil {
		middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", rcid, "CORE-MICRO", nil, respErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", rcid, "CORE-MICRO", nil, readErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := []response.SuperSaving{}

	unmErr := json.Unmarshal(body, &result)

	if unmErr != nil {
		middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", rcid, "CORE-MICRO", nil, unmErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", rcid, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer 		rebago
// @summary 	  	CORE-TRANSACTION-CHECKER
// @Description	  	Used to check if the transaction is posted in CORE
// @Tags		  	CORE MICRO
// @Accept		  	json
// @Produce		  	json
// @Param       	prNumber body request.PRNumber true  "Reference Number"
// @Success		  	200 {object} response.ResponseModelwoData
// @Failure		  	400 {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/callBackReference [post]
func CallBackReference(c *fiber.Ctx) error {
	prNumber := request.PRNumber{}

	if parsErr := c.BodyParser(&prNumber); parsErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, parsErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	marshallReq, marshallErr := json.Marshal(prNumber)

	fmt.Println("Marshall:", marshallErr)
	if marshallErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, marshallErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", "application/json; charset=utf-8", bytes.NewBuffer(marshallReq))

	if respErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, respErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, readErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := response.ResponseModelwoData{}

	unmErr := json.Unmarshal(body, &result)

	if unmErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, unmErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer	 Geromme
// @Summary      CORE-SAVINGS-INFORMATION
// @Description  Used to get the details of savings of NGO clients
// @Tags         CORE MICRO
// @Accept       json
// @Produce      json
// @Param        acc body request.Acc true "Account Number"
// @Success      200  {object} response.CustSavingInfo
// @Failure      400  {object} response.ResponseModel
// @Router       /api/public/v1/coremicro/custSavingInfo [post]
func CustSavingInfo(c *fiber.Ctx) error {
	acc := request.Acc{}

	if parsErr := c.BodyParser(&acc); parsErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, parsErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(acc)

	if marshallErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, marshallErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingInfo", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	if respErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, respErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, readErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := []response.CustSavingInfo{}
	unmErr := json.Unmarshal(body, &result)

	if unmErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, unmErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingInfo", acc, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer		Geromme
// @Summary			CORE-CLIENT-SAVINGS-LIST
// @Description 	Used to get the list of savings of NGO clients
// @Tags        	CORE MICRO
// @Accept      	json
// @Produce     	json
// @Param       	cid body request.CID true  "Client ID"
// @Success     	200  {object} response.CustSavingsLists
// @Failure      	400  {object} response.ResponseModel
// @Router      	/api/public/v1/coremicro/customerSavingsList/ [post]
func CustSavings(c *fiber.Ctx) error {
	cid := request.CID{}

	if parsErr := c.BodyParser(&cid); parsErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, parsErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(cid)

	if marshallErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, marshallErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingsList", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))

	if respErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, respErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, readErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := []response.CustSavingsLists{}
	unmErr := json.Unmarshal(body, &result)

	if unmErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, unmErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/custSavingsList", cid, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer		Geromme
// @Summary			CORE-DUE-GENERATOR
// @Description		Used to get the weekly due of the clients
// @Tags			CORE MICRO
// @Accept			json
// @Produce			json
// @Param       	cid body request.CID true  "Client ID"
// @Success			200  {object} response.GenerateCol
// @Failure      	400  {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/generateColCid/ [post]
func GenerateColCid(c *fiber.Ctx) error {
	cid := request.CID{}

	if parsErr := c.BodyParser(&cid); parsErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, parsErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(cid)

	if marshallErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, marshallErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/generateColShtperCID", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if respErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, respErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, readErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := []response.GenerateCol{}
	unmErr := json.Unmarshal(body, &result)

	if unmErr != nil {
		// middleware.SystemLoggerErrorAPI("https://cmfstest.cardmri.com/CoreMFS/API/k2cCallBackRef", prNumber, "CORE-MICRO", nil, unmErr, "localhost:8000")
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/generateColShtperCID", cid, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer		Geromme
// @Summary			CORE-LOAN-INFORMATION
// @Description		Used to get the information of loan and amortization
// @Tags			CORE MICRO
// @Accept			json
// @Produce			json
// @Param       	acc body request.Acc true  "Account Number"
// @Success			200  {object} response.LoanInfo
// @Failure      	400  {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/loanInfo [post]
func LoanInfo(c *fiber.Ctx) error {
	acc := request.Acc{}
	if parsErr := c.BodyParser(&acc); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(acc)
	if marshallErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/LoanInfo", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if respErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := []response.LoanInfo{}
	unmErr := json.Unmarshal(body, &result)

	if unmErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/LoanInfo", acc, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer		Geromme
// @Summary			CORE-LOAN-INFORMATION
// @Description		Used to get the information of loan and amortization
// @Tags			CORE MICRO
// @Accept			json
// @Produce			json
// @Param			paymentRequest body request.MultiplePaymentRequest true "Payment Request"
// @Success			200  {object} response.ResponseModelwoData
// @Failure      	400  {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/multiplePayment [post]
func MultiplePayment(c *fiber.Ctx) error {
	paymentRequest := request.MultiplePaymentRequest{}
	if parsErr := c.BodyParser(&paymentRequest); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(paymentRequest)
	if marshallErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreMFS/API/multiplePayment", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if respErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := response.ResponseModelwoData{}
	unmErr := json.Unmarshal(body, &result)
	if unmErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreMFS/API/multiplePayment", paymentRequest, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer		Geromme
// @Summary			CORE-OLD-LOAN-PAYMENT
// @Description		Old API used to pay loan in CORE
// @Tags			CORE MICRO
// @Accept			json
// @Produce			json
// @Param			paymentRequest body request.OpenPaymentRequest true "Search"
// @Success			200  {object} response.OpenPaymentResponse
// @Failure      	400  {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/openPaymentTransaction [post]
func OpenPayment(c *fiber.Ctx) error {

	openPayment := request.OpenPaymentRequest{}
	if parsErr := c.BodyParser(&openPayment); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(openPayment)
	if marshallErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/OpenPaymentTransaction", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if respErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := response.OpenPaymentResponse{}
	unmErr := json.Unmarshal(body, &result)
	if unmErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/OpenPaymentTransaction", openPayment, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer		Geromme
// @Summary			CORE-CLIENT-INFORMATION
// @Description		Used to get the details of NGO clients
// @Tags			CORE MICRO
// @Accept			json
// @Produce			json
// @Param       	cid body request.CID true  "Client ID"
// @Success			200  {object} response.CustomerInformation
// @Failure      	400  {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/searchCustomerCID [post]
func SearchCustomerCID(c *fiber.Ctx) error {
	cid := request.CID{}
	if parsErr := c.BodyParser(&cid); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(cid)
	if marshallErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/mobile/api/v1/SearchCustomerCID", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if respErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := response.CustomerInformation{}
	unmErr := json.Unmarshal(body, &result)
	if unmErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/mobile/api/v1/SearchCustomerCID", cid, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)
}

// Developer		Geromme
// @Summary			CORE-CLIENT-INFORMATION
// @Description		Used to get the details of NGO clients
// @Tags			CORE MICRO
// @Accept       	json
// @Produce      	json
// @Param       	cid body request.CID true  "Client ID"
// @Success      	200  {object} response.LoanList
// @Failure      	400  {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/searchLoanList [post]
func SearchLoanList(c *fiber.Ctx) error {
	cid := request.CID{}
	if parsErr := c.BodyParser(&cid); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	jsonReq, marshallErr := json.Marshal(cid)
	if marshallErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/searchLoanList", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if respErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    respErr,
		})
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Reading error",
			Data:    readErr,
		})
	}

	result := []response.LoanList{}
	unmErr := json.Unmarshal(body, &result)
	if unmErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	middleware.SystemLoggerAPI("https://cmfstest.cardmri.com/CoreAccounts/API/mobile/api/v1/searchLoanList", cid, "CORE-MICRO", resp, result, "localhost:8000")
	return c.JSON(result)

}
