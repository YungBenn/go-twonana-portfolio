package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "version": "1.0",
        "title": "Twonana Portfolio Go",
        "contact": {}
    },
    "host": "127.0.0.1:3000",
    "basePath": "/",
    "securityDefinitions": {},
    "schemes": [
        "http"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "paths": {
        "/api/v1/nft": {
            "post": {
                "summary": "Create NFT",
                "tags": [
                    "Nft"
                ],
                "operationId": "CreateNFT",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "parameters": [
                    {
                        "name": "title",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "image_url",
                        "in": "formData",
                        "required": true,
                        "type": "file",
                        "format": "file",
                        "description": ""
                    },
                    {
                        "name": "category",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "description",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "software",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "size",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "nft_format",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "marketplace_url",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "countdown_days",
                        "in": "formData",
                        "required": true,
                        "type": "integer",
                        "format": "int32",
                        "description": ""
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/CreateNFT"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 201,
                                    "message": "Created"
                                },
                                "data": {
                                    "_id": "64ce2cc6bddba9def8818e37",
                                    "title": "baby don't hurt me",
                                    "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1691233478/twonana%20nft/1e426798-0300-4bea-b2d3-58ade364fdac.png",
                                    "category": "dewa",
                                    "description": "blablabla",
                                    "software": "Adobe Illustrator",
                                    "size": "1000px x 1000px",
                                    "nft_format": "png",
                                    "marketplace_url": "url",
                                    "countdown_days": 2
                                }
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:04:38 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "405"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            },
            "get": {
                "summary": "Get List NFT",
                "tags": [
                    "Nft"
                ],
                "operationId": "GetListNFT",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetListNFT"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 200,
                                    "message": "OK"
                                },
                                "data": [
                                    {
                                        "_id": "64c8d83b0d878f31dc1467e4",
                                        "title": "ganjar",
                                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                                        "category": "PDIP",
                                        "description": "blablabla",
                                        "software": "Adobe Illustrator",
                                        "size": "1000px x 1000px",
                                        "nft_format": "png",
                                        "marketplace_url": "url",
                                        "countdown_days": 2
                                    },
                                    {
                                        "_id": "64c8d9840d878f31dc1467e5",
                                        "title": "baby don't hurt me",
                                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884484/twonana%20nft/818447a9-84e9-4045-8e75-11cc2b72af30.png",
                                        "category": "dewa",
                                        "description": "blablabla",
                                        "software": "Adobe Illustrator",
                                        "size": "1000px x 1000px",
                                        "nft_format": "png",
                                        "marketplace_url": "url",
                                        "countdown_days": 2
                                    },
                                    {
                                        "_id": "64ce2cc6bddba9def8818e37",
                                        "title": "baby don't hurt me",
                                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1691233478/twonana%20nft/1e426798-0300-4bea-b2d3-58ade364fdac.png",
                                        "category": "dewa",
                                        "description": "blablabla",
                                        "software": "Adobe Illustrator",
                                        "size": "1000px x 1000px",
                                        "nft_format": "png",
                                        "marketplace_url": "url",
                                        "countdown_days": 2
                                    }
                                ]
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:04:55 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "1104"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/nft/category": {
            "get": {
                "summary": "Get List NFT Category",
                "tags": [
                    "Nft"
                ],
                "operationId": "GetListNFTCategory",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetListNFTCategory"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 200,
                                    "message": "OK"
                                },
                                "data": [
                                    "PDIP",
                                    "dewa"
                                ]
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:05:02 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "59"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/nft/64c8d83b0d878f31dc1467e4": {
            "get": {
                "summary": "Get NFT By id",
                "tags": [
                    "Nft"
                ],
                "operationId": "GetNFTByid",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetNFTByid"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 200,
                                    "message": "OK"
                                },
                                "data": {
                                    "_id": "64c8d83b0d878f31dc1467e4",
                                    "title": "ganjar",
                                    "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                                    "category": "PDIP",
                                    "description": "blablabla",
                                    "software": "Adobe Illustrator",
                                    "size": "1000px x 1000px",
                                    "nft_format": "png",
                                    "marketplace_url": "url",
                                    "countdown_days": 2
                                }
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:05:24 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "388"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            },
            "put": {
                "summary": "Update NFT",
                "tags": [
                    "Nft"
                ],
                "operationId": "UpdateNFT",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "parameters": [
                    {
                        "name": "title",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "image_url",
                        "in": "formData",
                        "required": true,
                        "type": "file",
                        "format": "file",
                        "description": ""
                    },
                    {
                        "name": "category",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "description",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "software",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "size",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "nft_format",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "marketplace_url",
                        "in": "formData",
                        "required": true,
                        "type": "string",
                        "description": ""
                    },
                    {
                        "name": "countdown_days",
                        "in": "formData",
                        "required": true,
                        "type": "integer",
                        "format": "int32",
                        "description": ""
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/UpdateNFT"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 200,
                                    "message": "OK"
                                },
                                "data": {
                                    "_id": "64c8d83b0d878f31dc1467e4",
                                    "title": "ganjar",
                                    "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                                    "category": "PDIP",
                                    "description": "blablabla",
                                    "software": "Adobe Illustrator",
                                    "size": "1000px x 1000px",
                                    "nft_format": "png",
                                    "marketplace_url": "url",
                                    "countdown_days": 2
                                }
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:05:32 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "388"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/nft/64ce2cc6bddba9def8818e37": {
            "delete": {
                "summary": "Delete NFT",
                "tags": [
                    "Nft"
                ],
                "operationId": "DeleteNFT",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/DeleteNFT"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 200,
                                    "message": "OK"
                                },
                                "data": "Deleted"
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:05:50 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "53"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "summary": "Register",
                "tags": [
                    "Auth"
                ],
                "operationId": "Register",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/Register"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 201,
                                    "message": "Created"
                                },
                                "data": {
                                    "_id": "64ce2c9cbddba9def8818e36",
                                    "username": "twonana",
                                    "password": "$2a$10$Z..PJ8EJQRvPeWlYMlAcmultq2SJGYml4FxBx1Cw1vP0Nusu7ku.y"
                                }
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:03:55 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "178"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "summary": "Login",
                "tags": [
                    "Auth"
                ],
                "operationId": "Login",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Login"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 200,
                                    "message": "OK"
                                },
                                "data": "Login Success"
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:03:36 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "59"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            },
                            "Set-Cookie": {
                                "type": "string",
                                "default": "session_id=d190a577-351a-4bf3-9e5f-b37e2c48cbab; max-age=7200; path=/; SameSite=Lax"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "delete": {
                "summary": "Logout",
                "tags": [
                    "Auth"
                ],
                "operationId": "Logout",
                "deprecated": false,
                "produces": [
                    "application/json"
                ],
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Logout"
                        },
                        "examples": {
                            "application/json": {
                                "meta": {
                                    "code": 200,
                                    "message": "OK"
                                },
                                "data": "Logout Success"
                            }
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 11:04:08 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "60"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            },
                            "Set-Cookie": {
                                "type": "string",
                                "default": "session_id=; expires=Sat, 05 Aug 2023 11:03:08 GMT; path=/; SameSite=Lax"
                            }
                        }
                    }
                }
            }
        },
        "/": {
            "get": {
                "summary": "Health Check",
                "tags": [
                    "Misc"
                ],
                "operationId": "HealthCheck",
                "deprecated": false,
                "produces": [
                    "text/plain; charset=utf-8"
                ],
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string",
                            "default": ""
                        },
                        "examples": {
                            "text/plain; charset=utf-8": "OK!"
                        },
                        "headers": {
                            "Date": {
                                "type": "string",
                                "default": "Sat, 05 Aug 2023 09:22:59 GMT"
                            },
                            "Content-Length": {
                                "type": "string",
                                "default": "3"
                            },
                            "Vary": {
                                "type": "string",
                                "default": "Origin"
                            },
                            "Access-Control-Allow-Origin": {
                                "type": "string",
                                "default": "*"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateNFT": {
            "title": "CreateNFT",
            "example": {
                "meta": {
                    "code": 201,
                    "message": "Created"
                },
                "data": {
                    "_id": "64ce2cc6bddba9def8818e37",
                    "title": "baby don't hurt me",
                    "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1691233478/twonana%20nft/1e426798-0300-4bea-b2d3-58ade364fdac.png",
                    "category": "dewa",
                    "description": "blablabla",
                    "software": "Adobe Illustrator",
                    "size": "1000px x 1000px",
                    "nft_format": "png",
                    "marketplace_url": "url",
                    "countdown_days": 2
                }
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "$ref": "#/definitions/Data"
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "Meta": {
            "title": "Meta",
            "example": {
                "code": 201,
                "message": "Created"
            },
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "format": "int32"
                },
                "message": {
                    "type": "string"
                }
            },
            "required": [
                "code",
                "message"
            ]
        },
        "Data": {
            "title": "Data",
            "example": {
                "_id": "64ce2cc6bddba9def8818e37",
                "title": "baby don't hurt me",
                "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1691233478/twonana%20nft/1e426798-0300-4bea-b2d3-58ade364fdac.png",
                "category": "dewa",
                "description": "blablabla",
                "software": "Adobe Illustrator",
                "size": "1000px x 1000px",
                "nft_format": "png",
                "marketplace_url": "url",
                "countdown_days": 2
            },
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "software": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                },
                "nft_format": {
                    "type": "string"
                },
                "marketplace_url": {
                    "type": "string"
                },
                "countdown_days": {
                    "type": "integer",
                    "format": "int32"
                }
            },
            "required": [
                "_id",
                "title",
                "image_url",
                "category",
                "description",
                "software",
                "size",
                "nft_format",
                "marketplace_url",
                "countdown_days"
            ]
        },
        "GetListNFT": {
            "title": "GetListNFT",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": [
                    {
                        "_id": "64c8d83b0d878f31dc1467e4",
                        "title": "ganjar",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                        "category": "PDIP",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    },
                    {
                        "_id": "64c8d9840d878f31dc1467e5",
                        "title": "baby don't hurt me",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884484/twonana%20nft/818447a9-84e9-4045-8e75-11cc2b72af30.png",
                        "category": "dewa",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    },
                    {
                        "_id": "64ce2cc6bddba9def8818e37",
                        "title": "baby don't hurt me",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1691233478/twonana%20nft/1e426798-0300-4bea-b2d3-58ade364fdac.png",
                        "category": "dewa",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    }
                ]
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Data"
                    }
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "GetListNFTCategory": {
            "title": "GetListNFTCategory",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": [
                    "PDIP",
                    "dewa"
                ]
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "GetAllNFTByCategory": {
            "title": "GetAllNFTByCategory",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": [
                    {
                        "_id": "64c8d83b0d878f31dc1467e4",
                        "title": "ganjar",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                        "category": "PDIP",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    },
                    {
                        "_id": "64c8d9840d878f31dc1467e5",
                        "title": "baby don't hurt me",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884484/twonana%20nft/818447a9-84e9-4045-8e75-11cc2b72af30.png",
                        "category": "dewa",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    },
                    {
                        "_id": "64ce2cc6bddba9def8818e37",
                        "title": "baby don't hurt me",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1691233478/twonana%20nft/1e426798-0300-4bea-b2d3-58ade364fdac.png",
                        "category": "dewa",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    }
                ]
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Data"
                    }
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "GetNFTByTitle": {
            "title": "GetNFTByTitle",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": [
                    {
                        "_id": "64c8d83b0d878f31dc1467e4",
                        "title": "ganjar",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                        "category": "PDIP",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    },
                    {
                        "_id": "64c8d9840d878f31dc1467e5",
                        "title": "baby don't hurt me",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884484/twonana%20nft/818447a9-84e9-4045-8e75-11cc2b72af30.png",
                        "category": "dewa",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    },
                    {
                        "_id": "64ce2cc6bddba9def8818e37",
                        "title": "baby don't hurt me",
                        "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1691233478/twonana%20nft/1e426798-0300-4bea-b2d3-58ade364fdac.png",
                        "category": "dewa",
                        "description": "blablabla",
                        "software": "Adobe Illustrator",
                        "size": "1000px x 1000px",
                        "nft_format": "png",
                        "marketplace_url": "url",
                        "countdown_days": 2
                    }
                ]
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Data"
                    }
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "GetNFTByid": {
            "title": "GetNFTByid",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": {
                    "_id": "64c8d83b0d878f31dc1467e4",
                    "title": "ganjar",
                    "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                    "category": "PDIP",
                    "description": "blablabla",
                    "software": "Adobe Illustrator",
                    "size": "1000px x 1000px",
                    "nft_format": "png",
                    "marketplace_url": "url",
                    "countdown_days": 2
                }
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "$ref": "#/definitions/Data"
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "UpdateNFT": {
            "title": "UpdateNFT",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": {
                    "_id": "64c8d83b0d878f31dc1467e4",
                    "title": "ganjar",
                    "image_url": "https://res.cloudinary.com/dypwg6tc4/image/upload/v1690884155/twonana%20nft/3e05368d-c8e1-410e-bc60-28c2a96d944f.jpg",
                    "category": "PDIP",
                    "description": "blablabla",
                    "software": "Adobe Illustrator",
                    "size": "1000px x 1000px",
                    "nft_format": "png",
                    "marketplace_url": "url",
                    "countdown_days": 2
                }
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "$ref": "#/definitions/Data"
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "DeleteNFT": {
            "title": "DeleteNFT",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": "Deleted"
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "type": "string"
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "RegisterRequest": {
            "title": "RegisterRequest",
            "example": {
                "username": "twonana",
                "password": "twonana"
            },
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            },
            "required": [
                "username",
                "password"
            ]
        },
        "Register": {
            "title": "Register",
            "example": {
                "meta": {
                    "code": 201,
                    "message": "Created"
                },
                "data": {
                    "_id": "64ce2c9cbddba9def8818e36",
                    "username": "twonana",
                    "password": "$2a$10$Z..PJ8EJQRvPeWlYMlAcmultq2SJGYml4FxBx1Cw1vP0Nusu7ku.y"
                }
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "$ref": "#/definitions/Data6"
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "Data6": {
            "title": "Data6",
            "example": {
                "_id": "64ce2c9cbddba9def8818e36",
                "username": "twonana",
                "password": "$2a$10$Z..PJ8EJQRvPeWlYMlAcmultq2SJGYml4FxBx1Cw1vP0Nusu7ku.y"
            },
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            },
            "required": [
                "_id",
                "username",
                "password"
            ]
        },
        "LoginRequest": {
            "title": "LoginRequest",
            "example": {
                "username": "adminruben",
                "password": "adminruben"
            },
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            },
            "required": [
                "username",
                "password"
            ]
        },
        "Login": {
            "title": "Login",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": "Login Success"
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "type": "string"
                }
            },
            "required": [
                "meta",
                "data"
            ]
        },
        "Logout": {
            "title": "Logout",
            "example": {
                "meta": {
                    "code": 200,
                    "message": "OK"
                },
                "data": "Logout Success"
            },
            "type": "object",
            "properties": {
                "meta": {
                    "$ref": "#/definitions/Meta"
                },
                "data": {
                    "type": "string"
                }
            },
            "required": [
                "meta",
                "data"
            ]
        }
    },
    "tags": [
        {
            "name": "Nft"
        },
        {
            "name": "Auth"
        },
        {
            "name": "Misc",
            "description": ""
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Twonana Portfolio API",
	Description:      "Twonana Portfolio Website documentation with auth session, written by Ruben.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
