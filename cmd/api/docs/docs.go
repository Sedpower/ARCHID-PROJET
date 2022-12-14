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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
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
        "/api/aeroports": {
            "get": {
                "description": "retourne au format JSON tous les aeroports",
                "summary": "récupérer tous les aeroports",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/allMesure/{iata}/{date}": {
            "get": {
                "description": "get basic",
                "summary": "retourne la moyenne des mesures sur une journée d'un aeroport",
                "parameters": [
                    {
                        "type": "string",
                        "description": "aeroport Name",
                        "name": "iata",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start (format: YYYY-MM-DD)",
                        "name": "date",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Moyenne_Data_Day"
                            }
                        }
                    }
                }
            }
        },
        "/api/mesures/{iata}/{start}/{end}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "aeroport Name",
                        "name": "iata",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start (format: YYYY-MM-DD-HH)",
                        "name": "start",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End (format: YYYY-MM-DD-HH)",
                        "name": "end",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.MeasuresResultat"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.DayMeasurement": {
            "type": "object",
            "properties": {
                "HeureMesure": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.HourMeasurement"
                    }
                },
                "jour": {
                    "type": "string"
                }
            }
        },
        "main.HourMeasurement": {
            "type": "object",
            "properties": {
                "Mesure": {
                    "$ref": "#/definitions/main.Measurement"
                },
                "heure": {
                    "type": "string"
                }
            }
        },
        "main.Measurement": {
            "type": "object",
            "properties": {
                "Value": {
                    "type": "number"
                },
                "idCapteur": {
                    "type": "string"
                }
            }
        },
        "main.MeasuresResultat": {
            "type": "object",
            "properties": {
                "pressions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DayMeasurement"
                    }
                },
                "temperatures": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DayMeasurement"
                    }
                },
                "vitesseVents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.DayMeasurement"
                    }
                }
            }
        },
        "main.Moyenne_Data_Day": {
            "type": "object",
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Pressure": {
                    "type": "string"
                },
                "Temperature": {
                    "type": "string"
                },
                "Wind_speed": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Aeroport API",
	Description:      "Cette API vous permet d'effectuer deux get",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
