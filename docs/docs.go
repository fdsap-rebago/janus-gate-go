// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "FDSAP Support"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/public/v1/coremicro/getSavingsForSuperApp": {
            "post": {
                "description": "Super Saving",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "COREMICRO"
                ],
                "summary": "Super Saving",
                "parameters": [
                    {
                        "description": "Client ID",
                        "name": "cid",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SuperSaving"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CID": {
            "type": "object",
            "properties": {
                "cid": {
                    "description": "Client ID",
                    "type": "integer",
                    "example": 10724585
                }
            }
        },
        "response.ResponseModel": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "retCode": {
                    "type": "string"
                }
            }
        },
        "response.SuperSaving": {
            "type": "object",
            "properties": {
                "acc": {
                    "description": "Client Account",
                    "type": "string"
                },
                "balance": {
                    "description": "Account Balance",
                    "type": "number"
                },
                "centerCode": {
                    "description": "Center Code",
                    "type": "integer"
                },
                "centerName": {
                    "description": "Center Name",
                    "type": "string"
                },
                "cid": {
                    "description": "Client ID",
                    "type": "integer"
                },
                "fullname": {
                    "description": "Client Fullname",
                    "type": "string"
                },
                "unitCode": {
                    "description": "Unit Code",
                    "type": "integer"
                },
                "unitName": {
                    "description": "Unit Name",
                    "type": "string"
                },
                "withdrawable": {
                    "description": "Withdrawable",
                    "type": "number"
                },
                "withdrawalAmount": {
                    "description": "Withdrawal Amount",
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "JANUS V2",
	Description:      "JANUS Swagger JanusGate",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}