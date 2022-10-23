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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/balance/get_balance/{user_id}": {
            "get": {
                "description": "Возвращает баланс пользователя по его id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Получение баланса пользователя",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id пользователя",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.Balance"
                        }
                    }
                }
            }
        },
        "/balance/up_balance/{user_id}": {
            "post": {
                "description": "Пополняет баланс пользователя или создаёт его при первом пополнении",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Пополнение или инициализация баланса",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id пользователя",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.BalanceReplenishment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "avito2022_internal_app_api.Balance": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                }
            }
        },
        "avito2022_internal_app_api.BalanceReplenishment": {
            "type": "object",
            "properties": {
                "replenishment": {
                    "type": "number"
                }
            }
        },
        "internal_app_api.Balance": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                }
            }
        },
        "internal_app_api.BalanceReplenishment": {
            "type": "object",
            "properties": {
                "replenishment": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}