// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/devices": {
            "get": {
                "description": "Get a list of all devices in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "List all devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/device.DTO"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new device in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Create a new device",
                "parameters": [
                    {
                        "description": "Create device request object",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/device.CreateDeviceRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/device.DTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/err.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/devices/brand/{brand}": {
            "get": {
                "description": "Get all devices from a specific brand",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Find devices by brand",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device brand",
                        "name": "brand",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/device.DTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/devices/state/{state}": {
            "get": {
                "description": "Get all devices with a specific state",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Find devices by state",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device state",
                        "name": "state",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/device.DTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/devices/{id}": {
            "get": {
                "description": "Get a single device by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Get device by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/device.DTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a device by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Delete a device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an existing device by its ID, only devices that are not in the\nstate 'in_use' can be updated.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Update a device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated device request object",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/device.UpdateDeviceRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/err.Errors"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/err.Error"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Endpoint to perform a health check on the system",
                "tags": [
                    "Health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "device.CreateDeviceRequest": {
            "type": "object",
            "required": [
                "brand",
                "name",
                "state"
            ],
            "properties": {
                "brand": {
                    "type": "string",
                    "maxLength": 255
                },
                "name": {
                    "type": "string",
                    "maxLength": 255
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "available",
                        "in_use",
                        "inactive"
                    ]
                }
            }
        },
        "device.DTO": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "device.UpdateDeviceRequest": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "type": "string",
                    "enum": [
                        "available",
                        "in_use",
                        "inactive"
                    ]
                }
            }
        },
        "err.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "err.Errors": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Device Manager API",
	Description:      "API service for managing devices",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
