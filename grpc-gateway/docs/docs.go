// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "s.heidar@jabama.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/user-service/v1-login": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "V1Login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user_service"
                ],
                "summary": "V1Login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_service.V1LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.userServiceV1LoginResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.userServiceV1LoginResponseDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "payload": {
                    "$ref": "#/definitions/user_service.V1LoginResponse"
                }
            }
        },
        "user_service.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "user_service.V1LoginRequest": {
            "type": "object",
            "required": [
                "Name",
                "Password"
            ],
            "properties": {
                "Name": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "string"
                },
                "Password": {
                    "description": "@gotags: validate:\"required\"",
                    "type": "string"
                }
            }
        },
        "user_service.V1LoginResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/user_service.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "gRPC Gateway",
	Description:      "This is gRPC Gateway",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
