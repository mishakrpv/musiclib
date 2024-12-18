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
            "email": "mishavkrpv@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/songs": {
            "get": {
                "description": "get all songs matching filters from the library",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "query"
                ],
                "summary": "Songs",
                "parameters": [
                    {
                        "maxLength": 255,
                        "type": "string",
                        "description": "search by group",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "maxLength": 255,
                        "type": "string",
                        "description": "search by song",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "maxLength": 10,
                        "type": "string",
                        "description": "search by date",
                        "name": "date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by text",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "maxLength": 255,
                        "type": "string",
                        "description": "search by link",
                        "name": "link",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 150,
                        "minimum": 3,
                        "type": "integer",
                        "description": "songs per page",
                        "name": "songs",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/song.Song"
                            }
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "command"
                ],
                "summary": "CreateSong",
                "parameters": [
                    {
                        "description": "song to create",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/song.Song"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/{song_id}": {
            "put": {
                "description": "update song",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "command"
                ],
                "summary": "UpdateSong",
                "parameters": [
                    {
                        "maxLength": 255,
                        "type": "string",
                        "description": "Song ID",
                        "name": "song_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "song to update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete song",
                "tags": [
                    "command"
                ],
                "summary": "DeleteSong",
                "parameters": [
                    {
                        "maxLength": 255,
                        "type": "string",
                        "description": "Song ID",
                        "name": "song_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/songs/{song_id}/lyrics": {
            "get": {
                "description": "get song's lyrics with pagination by verses",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "query"
                ],
                "summary": "Lyrics",
                "parameters": [
                    {
                        "maxLength": 255,
                        "type": "string",
                        "description": "Song ID",
                        "name": "song_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "verse number",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "verse",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "command.CreateRequest": {
            "type": "object",
            "required": [
                "group",
                "song"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "maxLength": 255
                },
                "song": {
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "command.UpdateRequest": {
            "type": "object",
            "required": [
                "group",
                "song"
            ],
            "properties": {
                "group": {
                    "type": "string",
                    "maxLength": 255
                },
                "link": {
                    "type": "string",
                    "maxLength": 255
                },
                "release_date": {
                    "type": "string",
                    "maxLength": 10
                },
                "song": {
                    "type": "string",
                    "maxLength": 255
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "song.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Musiclib API",
	Description:      "Effective Mobile test task",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
