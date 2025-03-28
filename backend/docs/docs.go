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
        "license": {
            "name": "tetn39"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/hello": {
            "get": {
                "description": "Hello Worldを返すだけのテスト用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hello World"
                ],
                "summary": "hello worldを返す",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HelloStruct"
                        }
                    }
                }
            }
        },
        "/musics": {
            "get": {
                "description": "全ての音楽データを取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Musics"
                ],
                "summary": "すべての音楽データを取得する",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SwaggerMusic"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "音楽データを登録する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Musics"
                ],
                "summary": "音楽データを登録する",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SwaggerMusic"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/musics/{id}": {
            "delete": {
                "description": "{id}の音楽データを削除する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Musics"
                ],
                "summary": "音楽データを削除する",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DeletedResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DeletedResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                }
            }
        },
        "models.HelloStruct": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.SwaggerMusic": {
            "type": "object",
            "properties": {
                "acousticness": {
                    "type": "number"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "danceability": {
                    "type": "number"
                },
                "deleted_at": {
                    "type": "string"
                },
                "energy": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "instrumentalness": {
                    "type": "number"
                },
                "liveness": {
                    "type": "number"
                },
                "loudness": {
                    "type": "number"
                },
                "mode": {
                    "type": "number"
                },
                "speechiness": {
                    "type": "number"
                },
                "tempo": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "valence": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "MelodyStatusのAPI一覧",
	Description:      "MelodyStatusのAPI一覧だよ",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
