// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Skyler"
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
        "/Device": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Get Account by index",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "index",
                        "name": "index",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"OK\"}",
                        "schema": {
                            "$ref": "#/definitions/main.account"
                        }
                    },
                    "400": {
                        "description": "{\"msg\":\"fail\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Add Account in Slice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"OK\"}",
                        "schema": {
                            "$ref": "#/definitions/main.account"
                        }
                    },
                    "400": {
                        "description": "{\"msg\":\"fail\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Delete Account by index",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "auth",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "index",
                        "name": "index",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"OK\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"msg\":\"fail\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.account": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{"http,https"},
	Title:       "swagger test",
	Description: "swagger test example",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
