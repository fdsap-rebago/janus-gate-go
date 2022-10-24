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
// @summary 	  	Super Saving
// @Description	  	Super Saving
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
// @summary 	  	Call Back Reference
// @Description	  	Call Back Reference
// @Tags		  	CORE MICRO
// @Accept		  	json
// @Produce		  	json
// @Param       	prNumber body request.PRNumber true  "Reference Number"
// @Success		  	200 {object} response.ResponseModel
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

	result := response.ReferenceCallBack{}

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
// @Summary      Customer Savings Information
// @Description  Customer Savings Information
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
// @Summary			Customer Savings List
// @Description 	Customer Savings List
// @Tags        	CORE MICRO
// @Accept      	json
// @Produce     	json
// @Param       	cid body request.CID true  "Client ID"
// @Success     	200  {object} response.CustSavingsLists
// @Failure      	400  {object} response.ResponseModel
// @Router      	/api/public/v1/coremicro/customerSavings/ [post]
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
