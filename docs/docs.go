// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/signin": {
            "post": {
                "description": "Authenticate user and return JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Sign in credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.SignInResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Register a new user with username, password and type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Sign up request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.SignUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "409": {
                        "description": "Username already exists",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "503": {
                        "description": "Database connection error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/avatars": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get list of all available avatars",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "avatars"
                ],
                "summary": "Get available avatars",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.GetAvatarsResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/element": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create an element with specified properties",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "elements"
                ],
                "summary": "Create a new element",
                "parameters": [
                    {
                        "description": "Element creation request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateElementRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateElementResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/element/all": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a list of all available elements",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "elements"
                ],
                "summary": "Get all elements",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.GetAllElementsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/element/{elementId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an element's image URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "elements"
                ],
                "summary": "Update an element",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Element ID",
                        "name": "elementId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Element update request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UpdateElementRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated"
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "404": {
                        "description": "Element not found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/space": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a space with specified name, dimensions and map ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spaces"
                ],
                "summary": "Create a new space",
                "parameters": [
                    {
                        "description": "Space creation request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateSpaceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.CreateSpaceResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/space/all": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all spaces for the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spaces"
                ],
                "summary": "Get all spaces",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.GetAllSpacesResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/space/{spaceId}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a space by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "spaces"
                ],
                "summary": "Delete a space",
                "parameters": [
                    {
                        "type": "string",
                        "example": "123e4567-e89b-12d3-a456-426614174000",
                        "description": "Space ID",
                        "name": "spaceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted"
                    },
                    "400": {
                        "description": "Invalid space ID",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "404": {
                        "description": "Space not found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/users/metadata": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user's avatar and other metadata",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user metadata",
                "parameters": [
                    {
                        "description": "Update metadata request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UpdateMetadataRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UpdateMetadataResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "401": {
                        "description": "Invalid token",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        },
        "/users/metadata/bulk": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get metadata for multiple users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users metadata bulk",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User IDs array format: [uuid1,uuid2,uuid3]",
                        "name": "ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.GetUserMetadataBulkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrorStruct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Element": {
            "type": "object",
            "properties": {
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "static": {
                    "type": "boolean"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "schemas.AvatarResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schemas.CreateElementRequest": {
            "type": "object",
            "required": [
                "height",
                "imageURL",
                "static",
                "width"
            ],
            "properties": {
                "height": {
                    "type": "integer"
                },
                "imageURL": {
                    "type": "string"
                },
                "static": {
                    "type": "boolean"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "schemas.CreateElementResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "schemas.CreateSpaceRequest": {
            "type": "object",
            "required": [
                "mapId",
                "name"
            ],
            "properties": {
                "mapId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schemas.CreateSpaceResponse": {
            "type": "object",
            "properties": {
                "spaceId": {
                    "type": "string"
                }
            }
        },
        "schemas.GetAllElementsResponse": {
            "type": "object",
            "properties": {
                "elements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Element"
                    }
                }
            }
        },
        "schemas.GetAllSpacesResponse": {
            "type": "object",
            "properties": {
                "spaces": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.SpaceResponse"
                    }
                }
            }
        },
        "schemas.GetAvatarsResponse": {
            "type": "object",
            "properties": {
                "avatars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.AvatarResponse"
                    }
                }
            }
        },
        "schemas.GetUserMetadataBulkResponse": {
            "type": "object",
            "properties": {
                "avatars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.UserMetadataResponse"
                    }
                }
            }
        },
        "schemas.SignInRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.SignInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "schemas.SignUpRequest": {
            "type": "object",
            "required": [
                "password",
                "type",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "Admin",
                        "User"
                    ]
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "schemas.SignUpResponse": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "string"
                }
            }
        },
        "schemas.SpaceResponse": {
            "type": "object",
            "properties": {
                "Thumbnail": {
                    "type": "string"
                },
                "creatorID": {
                    "type": "string"
                },
                "dimensions": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schemas.UpdateElementRequest": {
            "type": "object",
            "required": [
                "imageUrl"
            ],
            "properties": {
                "imageUrl": {
                    "type": "string"
                }
            }
        },
        "schemas.UpdateMetadataRequest": {
            "type": "object",
            "required": [
                "avatarId"
            ],
            "properties": {
                "avatarId": {
                    "type": "string"
                }
            }
        },
        "schemas.UpdateMetadataResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "schemas.UserMetadataResponse": {
            "type": "object",
            "properties": {
                "imageUrl": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "utils.ErrorStruct": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
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
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "SPAC REST API",
	Description:      "A RESTful API service for SPAC",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
