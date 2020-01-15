// Code generated by go-swagger; DO NOT EDIT.

package rest_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "dtaservice.proto",
    "version": "version not set"
  },
  "paths": {
    "/v1/document/transform-pipe": {
      "post": {
        "tags": [
          "DTAServer"
        ],
        "operationId": "TransformPipe",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dtaserviceTransformPipeRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/dtaserviceTransformDocumentReply"
            }
          }
        }
      }
    },
    "/v1/dta/document/transform": {
      "post": {
        "tags": [
          "DTAServer"
        ],
        "summary": "Request to transform a plain text document",
        "operationId": "TransformDocument",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dtaserviceDocumentRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/dtaserviceTransformDocumentReply"
            }
          }
        }
      }
    },
    "/v1/dta/service/list": {
      "get": {
        "tags": [
          "DTAServer"
        ],
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/dtaserviceListServicesResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "dtaserviceDocumentRequest": {
      "type": "object",
      "title": "The request message containing the document to be transformed",
      "properties": {
        "document": {
          "type": "string",
          "format": "byte"
        },
        "file_name": {
          "type": "string"
        },
        "options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "service_name": {
          "type": "string"
        }
      }
    },
    "dtaserviceListServicesResponse": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "dtaserviceTransformDocumentReply": {
      "type": "object",
      "title": "The response message containing the transformed message",
      "properties": {
        "error": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "trans_document": {
          "type": "string",
          "format": "byte"
        },
        "trans_output": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "dtaserviceTransformPipeRequest": {
      "type": "object",
      "properties": {
        "pipe": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dtaserviceDocumentRequest"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "dtaservice.proto",
    "version": "version not set"
  },
  "paths": {
    "/v1/document/transform-pipe": {
      "post": {
        "tags": [
          "DTAServer"
        ],
        "operationId": "TransformPipe",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dtaserviceTransformPipeRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/dtaserviceTransformDocumentReply"
            }
          }
        }
      }
    },
    "/v1/dta/document/transform": {
      "post": {
        "tags": [
          "DTAServer"
        ],
        "summary": "Request to transform a plain text document",
        "operationId": "TransformDocument",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/dtaserviceDocumentRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/dtaserviceTransformDocumentReply"
            }
          }
        }
      }
    },
    "/v1/dta/service/list": {
      "get": {
        "tags": [
          "DTAServer"
        ],
        "operationId": "ListServices",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/dtaserviceListServicesResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "dtaserviceDocumentRequest": {
      "type": "object",
      "title": "The request message containing the document to be transformed",
      "properties": {
        "document": {
          "type": "string",
          "format": "byte"
        },
        "file_name": {
          "type": "string"
        },
        "options": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "service_name": {
          "type": "string"
        }
      }
    },
    "dtaserviceListServicesResponse": {
      "type": "object",
      "properties": {
        "services": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "dtaserviceTransformDocumentReply": {
      "type": "object",
      "title": "The response message containing the transformed message",
      "properties": {
        "error": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "trans_document": {
          "type": "string",
          "format": "byte"
        },
        "trans_output": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "dtaserviceTransformPipeRequest": {
      "type": "object",
      "properties": {
        "pipe": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/dtaserviceDocumentRequest"
          }
        }
      }
    }
  }
}`))
}
