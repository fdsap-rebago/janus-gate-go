package coremicro

import (
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

// JANUS GATE - Super Saving
// @summary 	  	Super Saving
// @Description	  	Super Saving
// @Tags		  	COREMICRO
// @Accept		  	json
// @Produce		  	json
// @Param       	cid body request.CID true  "Client ID"
// @Success		  	200 {object} response.SuperSaving
// @Failure		  	400 {object} response.ResponseModel
// @Router			/api/public/v1/coremicro/getSavingsForSuperApp [post]
func SuperAppSaving(c *fiber.Ctx) error {
	rcid := request.CID{}

	if parsErr := c.BodyParser(&rcid); parsErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Bad request",
			Data:    parsErr,
		})
	}

	marshallReq, marshallErr := json.Marshal(rcid.Cid)

	fmt.Println("Marshall:", marshallErr)
	if marshallErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Can't masrhall",
			Data:    marshallErr,
		})
	}

	resp, respErr := http.Post("https://cmfstest.cardmri.com/CoreAccounts/API/getSavingForSuperApp", "application/json; charset=utf-8", bytes.NewBuffer(marshallReq))

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

	result := []response.SuperSaving{}

	unmErr := json.Unmarshal(body, &result)

	if unmErr != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "Unmarshall error",
			Data:    unmErr,
		})
	}

	return c.JSON(result)
}
