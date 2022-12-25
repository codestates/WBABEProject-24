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
        "/orderer/menu/list": {
            "get": {
                "description": "Menu 객체 리스트를 반환하기 위한 기능",
                "produces": [
                    "application/json"
                ],
                "summary": "call GetMenuList, return the Menu object list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sort type [recommend|score|mostOrder|new]",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data: 메뉴 리스트",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.SuccessResultJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Menu"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/orderer/order": {
            "post": {
                "description": "주문 객체를 생성하기 위한 기능",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateOrder, Create the Order object and return result message by string.",
                "parameters": [
                    {
                        "description": "Order data",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data: 생성된 주문 일련번호",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.SuccessResultJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/orderer/order/list": {
            "get": {
                "description": "주문 상태에 해당하는 리스트를 반환하는 기능",
                "produces": [
                    "application/json"
                ],
                "summary": "call GetOrderList, return the Order object list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order status [active|deactive|complete|all]",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data: status에 해당하는 주문 리스트",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.SuccessResultJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Order"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/orderer/order/{seq}/{type}": {
            "put": {
                "description": "주문의 메뉴를 변경하기 위한 기능",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call ChangeOrderMenu, Change the Order menu and return result message by string.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order sequence number",
                        "name": "seq",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order change type [add|change]",
                        "name": "type",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "새로운 주문 일련번호",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.SuccessResultJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/orderer/review": {
            "post": {
                "description": "리뷰 객체를 생성하기 위한 기능",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateReview, Create the Review object and return result message by string.",
                "parameters": [
                    {
                        "description": "Review data",
                        "name": "review",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Review"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SuccessResultJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/orderer/review/list/{menu}": {
            "get": {
                "description": "메뉴에 해당하는 리뷰 리스트를 반환하기 위한 기능",
                "produces": [
                    "application/json"
                ],
                "summary": "call GetReviewList, return the List of reviews corresponding to the menu.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu name",
                        "name": "menu",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data: 메뉴에 해당하는 리뷰 리스트",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.SuccessResultJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Review"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/recipant/menu": {
            "post": {
                "description": "Menu 객체를 생성하기 위한 기능",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call CreateMenu, Create the Menu object and return result message by string.",
                "parameters": [
                    {
                        "description": "Menu data",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Menu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SuccessResultJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/recipant/menu/{name}": {
            "put": {
                "description": "Menu 객체를 업데이트하기 위한 기능",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "call UpdateMenu, Update the Menu object and return result message by string.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu name for update",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Menu data for update",
                        "name": "menu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Menu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SuccessResultJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            },
            "delete": {
                "description": "Menu 객체를 삭제하기 위한 기능",
                "produces": [
                    "application/json"
                ],
                "summary": "call DeleteMenu, Delete the Menu object and return result message by string.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Menu name for delete",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SuccessResultJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/recipant/order/list": {
            "get": {
                "description": "주문 상태에 해당하는 리스트를 반환하는 기능",
                "produces": [
                    "application/json"
                ],
                "summary": "call GetOrderList, return the Order object list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order status [active|deactive|complete|all]",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "data: status에 해당하는 주문 리스트",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/controller.SuccessResultJSON"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Order"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        },
        "/recipant/order/{seq}/{status}": {
            "put": {
                "description": "주문의 상태를 변경하기 위한 기능",
                "produces": [
                    "application/json"
                ],
                "summary": "call ChangeOrderStatus, Change the Order status and return result message by string.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order sequence number",
                        "name": "seq",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "status value to change [대기|주문|조리|배달|완료]",
                        "name": "status",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SuccessResultJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResultJSON"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ErrorResultJSON": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controller.SuccessResultJSON": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Menu": {
            "type": "object"
        },
        "model.Order": {
            "type": "object"
        },
        "model.Review": {
            "type": "object",
            "required": [
                "comment",
                "menuName",
                "orderSeq",
                "score"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                },
                "menuName": {
                    "type": "string"
                },
                "orderSeq": {
                    "type": "string"
                },
                "reviewNum": {
                    "type": "string"
                },
                "score": {
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
