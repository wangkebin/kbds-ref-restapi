// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
  "swagger": "2.0",
  "info": {
    "title": "kbds-ref-restapi",
    "version": "0.1.0"
  },
  "paths": {
    "/v1/delete/{fileid}": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "file to delete.",
            "name": "fileid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully deleted",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/duplicates": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "duplicates",
        "parameters": [
          {
            "description": "file info of interest",
            "name": "finfos",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/files"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully posted the list of fileinfo resources",
            "schema": {
              "$ref": "#/definitions/files"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/files": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "files",
        "parameters": [
          {
            "description": "list of file info to post",
            "name": "files",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/files"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully posted the list of fileinfo resources",
            "schema": {
              "type": "integer"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "deletefiles",
        "parameters": [
          {
            "description": "list of file info to delete",
            "name": "files",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/files"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully deleted the list of fileinfo resources",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/healthz": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Successfully checked health status",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/search": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "search",
        "parameters": [
          {
            "description": "part of file name/path you are searching for",
            "name": "search",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page": {
                  "type": "integer"
                },
                "pagesize": {
                  "type": "integer"
                },
                "search": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully searched the file list based on a search term",
            "schema": {
              "$ref": "#/definitions/files"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "file": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "count": {
          "type": "integer"
        },
        "ext": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "loc": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "size": {
          "type": "integer"
        }
      }
    },
    "files": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/file"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "kbds-ref-restapi",
    "version": "0.1.0"
  },
  "paths": {
    "/v1/delete/{fileid}": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "file to delete.",
            "name": "fileid",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully deleted",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/duplicates": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "duplicates",
        "parameters": [
          {
            "description": "file info of interest",
            "name": "finfos",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/files"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully posted the list of fileinfo resources",
            "schema": {
              "$ref": "#/definitions/files"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/files": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "files",
        "parameters": [
          {
            "description": "list of file info to post",
            "name": "files",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/files"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully posted the list of fileinfo resources",
            "schema": {
              "type": "integer"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "deletefiles",
        "parameters": [
          {
            "description": "list of file info to delete",
            "name": "files",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/files"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully deleted the list of fileinfo resources",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/healthz": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Successfully checked health status",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/search": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "operationId": "search",
        "parameters": [
          {
            "description": "part of file name/path you are searching for",
            "name": "search",
            "in": "body",
            "required": true,
            "schema": {
              "properties": {
                "page": {
                  "type": "integer"
                },
                "pagesize": {
                  "type": "integer"
                },
                "search": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully searched the file list based on a search term",
            "schema": {
              "$ref": "#/definitions/files"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "file": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "count": {
          "type": "integer"
        },
        "ext": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "loc": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "size": {
          "type": "integer"
        }
      }
    },
    "files": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/file"
      }
    }
  }
}`))
}
