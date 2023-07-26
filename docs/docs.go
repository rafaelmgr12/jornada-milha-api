// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Rafael Ribeiro",
            "email": "ribeirorafaelmatheus@gmail.com"
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
        "/api/v1/depoimentos": {
            "post": {
                "description": "Create a Testimonial",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "testimonials"
                ],
                "summary": "Create a Testimonial",
                "parameters": [
                    {
                        "description": "TestimonialCreateDTO",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/testimonials.TestimonialCreateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Testimonial"
                        }
                    }
                }
            }
        },
        "/api/v1/testimonials": {
            "get": {
                "description": "Get a list of all Testimonials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "testimonials"
                ],
                "summary": "Retrieve all Testimonials",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Testimonial"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update a specific Testimonial",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "testimonials"
                ],
                "summary": "Update a Testimonial",
                "parameters": [
                    {
                        "description": "TestimonialUpdateDTO",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/testimonials.TestimonialUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Testimonial"
                        }
                    }
                }
            }
        },
        "/api/v1/testimonials/{id}": {
            "delete": {
                "description": "Delete a specific Testimonial",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "testimonials"
                ],
                "summary": "Delete a Testimonial",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Testimonial ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        "entity.Testimonial": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "testimonial": {
                    "type": "string"
                }
            }
        },
        "testimonials.TestimonialCreateDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "testimonial": {
                    "type": "string"
                }
            }
        },
        "testimonials.TestimonialUpdateDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "testimonial": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Jorna Milha API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}