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
        "/accounting/report_link": {
            "post": {
                "description": "Метод по году и месяцу возвращает ссылку на скачивание отчёта",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounting"
                ],
                "summary": "Получение ссылки с отчётом за месяц",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/avito2022_internal_app_api.ReportRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/avito2022_internal_app_api.ReportLink"
                        }
                    }
                }
            }
        },
        "/balance/get_balance": {
            "post": {
                "description": "Возвращает баланс пользователя по его id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Получение баланса пользователя",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.UserID"
                        }
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
        "/balance/transfer_money": {
            "post": {
                "description": "Метод перевода денег от пользователя к пользователю",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Перевод денег",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.Transfer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/balance/up_balance": {
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
        },
        "/balance/withdraw_money": {
            "post": {
                "description": "Метод снимает указанное количество средств со счёта",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Снятие денег со счёта",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.DebitingMoney"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/sales/reserve_money": {
            "post": {
                "description": "Резервирование средств с основного баланса на отдельном счете",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sales"
                ],
                "summary": "Резервирование средств",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.Reserve"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/sales/return_money": {
            "post": {
                "description": "Разрезервирование средств и возвращение их на баланс пользователя при отмене операции",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sales"
                ],
                "summary": "Возвращение средств",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.Refund"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/sales/revenue_confirmation": {
            "post": {
                "description": "Списывает из резерва деньги, добавляет данные в отчет для бухгалтерии",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "sales"
                ],
                "summary": "Метод признания выручки",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_app_api.Confirmation"
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
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "avito2022_internal_app_api.Confirmation": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "order_id": {
                    "type": "integer"
                },
                "service_id": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "avito2022_internal_app_api.DebitingMoney": {
            "type": "object",
            "properties": {
                "debit_cost": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "avito2022_internal_app_api.Refund": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "integer"
                }
            }
        },
        "avito2022_internal_app_api.ReportLink": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                }
            }
        },
        "avito2022_internal_app_api.ReportRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                }
            }
        },
        "avito2022_internal_app_api.Reserve": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "order_id": {
                    "type": "integer"
                },
                "service_id": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "avito2022_internal_app_api.Transfer": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "recipient_id": {
                    "type": "integer"
                },
                "recipient_name": {
                    "type": "string"
                },
                "sender_id": {
                    "type": "integer"
                },
                "sender_name": {
                    "type": "string"
                }
            }
        },
        "avito2022_internal_app_api.UserID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
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
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "internal_app_api.Confirmation": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "order_id": {
                    "type": "integer"
                },
                "service_id": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "internal_app_api.DebitingMoney": {
            "type": "object",
            "properties": {
                "debit_cost": {
                    "type": "number"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "internal_app_api.Refund": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "integer"
                }
            }
        },
        "internal_app_api.ReportLink": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                }
            }
        },
        "internal_app_api.ReportRequest": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                }
            }
        },
        "internal_app_api.Reserve": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "order_id": {
                    "type": "integer"
                },
                "service_id": {
                    "type": "integer"
                },
                "service_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "internal_app_api.Transfer": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "recipient_id": {
                    "type": "integer"
                },
                "recipient_name": {
                    "type": "string"
                },
                "sender_id": {
                    "type": "integer"
                },
                "sender_name": {
                    "type": "string"
                }
            }
        },
        "internal_app_api.UserID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
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
