package jsonschema

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qri-io/jsonschema"
)

func Do2() {
	ctx := context.Background()
	var schemaData = []byte(`{
		"$schema": "http://json-schema.org/draft-04/schema#",
		"type": "array",
		"items": [
		  {
			"type": "object",
			"properties": {
			  "consumer": {
				"type": "object",
				"properties": {
				  "name": {
					"type": "string"
				  },
				  "config": {
					"type": "object",
					"additionalProperties": true
				  }
				},
				"required": [
				  "name",
				  "config"
				]
			  },
			  "processors": {
				"type": "array",
				"items": [
				  {
					"type": "object",
					"properties": {
					  "name": {
						"type": "string"
					  },
					  "transformer": {
						"type": "object",
						"properties": {
						  "type": {
							"type": "string"
						  },
						  "config": {
							"type": "object",
							"additionalProperties": true
						  }
						},
						"required": [
						  "type",
						  "config"
						]
					  },
					  "storage": {
						"type": "array",
						"items": [
						  {
							"type": "object",
							"properties": {
							  "type": {
								"type": "string"
							  },
							  "config": {
								"type": "object",
								"additionalProperties": true
							  }
							},
							"required": [
							  "type",
							  "config"
							]
						  }
						],
						"additionalItems": true
					  },
					  "content_type_filter": {
						"type": "array",
						"items": {},
						"additionalItems": true
					  },
					  "cache": {
						"type": "object",
						"properties": {
						  "type": {
							"type": "string"
						  },
						  "config": {
							"type": "object",
							"additionalProperties": true
						  }
						},
						"additionalProperties": true,
						"required": [
						  "type",
						  "config"
						]
					  }
					},
					"required": [
					  "name",
					  "transformer",
					  "storage",
					  "content_type_filter",
					  "cache"
					]
				  }
				],
				"additionalItems": true
			  }
			},
			"required": [
			  "consumer",
			  "processors"
			]
		  }
		],
		"additionalItems": true
	  }`)

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	var valid = []byte(`[
		{
			"consumer": {
				"name": "confluent_kafka_client",
				"config": {
					"max_retries": -1,
					"retry_delay": 5,
					"record_check_frequency": 15,
					"client_name": "default",
					"topic": "PROVIDE AN APP ID",
					"consumer_options": {
						"bootstrap.servers": "localhost:9092",
						"group.id": "01",
						"auto.offset.reset": "earliest",
						"enable.auto.commit": "false",
						"enable.partition.eof": "true",
						"session.timeout.ms": "60000",
						"heartbeat.interval.ms": "15000"
					}
				}
			},
			"processors": [
				{
					"name": "nuance-mix3",
					"transformer": {
						"type": "nuance-mix3",
						"config": {
							"kv_mappings": {
								"app_id": "OPTIONAL: PROVIDE A TARGET MIX 3 APP ID TO MAP KAFKA TOPIC TO"
							}
						}
					},
					"storage": [
						{
							"type": "filesystem",
							"config": {
								"path": "default-processor-logs/customerA-appname-dev",
								"queue": {
									"size": 100000,
									"workers": 15
								}
							}
						},
						{
							"type": "s3",
							"config": {
								"endpoint": "localhost:9000",
								"access_key_id": "minio",
								"secret_access_key": "minio123",
								"use_ssl": false,
								"queue": {
									"size": 100000,
									"workers": 15
								}
							}
						},
						{
							"type": "nuance-mix3",
							"config": {
								"credentials": {
									"auth_disabled": false,
									"token_url": "https://auth.crt.nuance.com/oauth2/token",
									"scope": "log",
									"client_id": "PROVIDE A CLIENT ID",
									"client_secret": "PROVIDE A CLIENT SECRET",
									"auth_timeout_ms": 5000
								},
								"api_url": "https://log.api.nuance.com/producers",
								"connect_timeout_ms": 5000,
								"rate_limit": 8,
								"max_retries": -1,
								"delay": 5,
								"threshold_bytes": 500000,
								"queue": {
									"size": 100000,
									"workers": 15
								},
								"cache": {
									"type": "redis",
									"config": {
										"addr": "localhost:6379",
										"password": "provide redis password",
										"db": 0,
										"default_ttl": 900
									}
								}
							}
						}
					],
					"cache": {
						"type": "local",
						"config": {
							"expiration": 2880,
							"cleanup_interval": 300
						}
					},
					"content_type_filter": [
									"application/x-nuance-nluaas-interpretation.v1+json",
									"application/x-nuance-asr-finalresultresponse.v2+json",
									"application/x-nuance-asr-partialresultresponse.v2+json",
									"application/x-nuance-asr-recognitioninitmessage.v2+json",
									"application/x-nuance-asr-statusresponse.v2+json",
									"application/x-nuance-asr-startspeechresponse.v2+json"
					]
				}
			]
		}
	]`)
	errs, err := rs.ValidateBytes(ctx, valid)
	if err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	}
}

func Do() {
	ctx := context.Background()
	var schemaData = []byte(`{
    "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Person",
    "type": "object",
    "properties": {
        "firstName": {
            "type": "string"
        },
        "lastName": {
            "type": "string"
        },
        "age": {
            "description": "Age in years",
            "type": "integer",
            "minimum": 0
        },
        "friends": {
          "type" : "array",
          "items" : { "title" : "REFERENCE", "$ref" : "#" }
        }
    },
    "required": ["firstName", "lastName"]
  }`)

	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	var valid = []byte(`{
    "firstName" : "George",
    "lastName" : "Michael"
    }`)
	errs, err := rs.ValidateBytes(ctx, valid)
	if err != nil {
		panic(err)
	}

	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	}

	var invalidPerson = []byte(`{
    "firstName" : "Prince"
    }`)

	errs, err = rs.ValidateBytes(ctx, invalidPerson)
	if err != nil {
		panic(err)
	}
	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	}

	var invalidFriend = []byte(`{
    "firstName" : "Jay",
    "lastName" : "Z",
    "friends" : [{
      "firstName" : "Nas"
      }]
    }`)
	errs, err = rs.ValidateBytes(ctx, invalidFriend)
	if err != nil {
		panic(err)
	}
	if len(errs) > 0 {
		fmt.Println(errs[0].Error())
	}
}
