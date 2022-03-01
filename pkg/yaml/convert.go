package yaml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
	yml "github.com/ghodss/yaml"
)

type jsonObj map[string]interface{}

func DoConvertJsonToYaml() {

	var jsonData = make(jsonObj)

	jsonData["hi"] = "there"

	data, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	yamlD, err := yml.JSONToYAML(data)
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	ioutil.WriteFile("ConvertedFromJson.yaml", yamlD, 0644)
}
func UnMarshallYamlToMap() {

	data, err := ioutil.ReadFile("ConvertedFromJson.yaml")
	if err != nil {
		log.Panic(err)
	}
	var jsonData jsonObj
	err = yml.Unmarshal(data, &jsonData)
	if err != nil {
		log.Panic(err)
	}

	if hi, ok := jsonData["hi"]; ok {
		log.Println(hi)
	}
}

func UnMarshalJsonYamlStruct() {
	type JsonYamlClass struct {
		Greeting string `json:"hi" yaml:"hi"`
	}
	jsonData, _ := ioutil.ReadFile("test.json")
	var jsonObj JsonYamlClass
	json.Unmarshal(jsonData, &jsonObj)
	fmt.Println(jsonObj.Greeting)

	yamlData, _ := ioutil.ReadFile("test.yaml")
	var yaml JsonYamlClass

	yml.Unmarshal(yamlData, &yaml)
	fmt.Println(yaml.Greeting)

}

func ConvertJsonConfigToYaml() {
	config := `[
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
							"type": "elasticsearch",
							"config": {
								"addresses": [
									"http://localhost:9200"
								],
								"workerpool": {
									"size": 1
								},
								"rate_limit": {
									"limit": 10,
									"burst": 1
								}
							}
						},
						{
							"type": "filesystem",
							"config": {
								"path": "default-processor-logs/customerA-appname-dev",
								"workerpool": {
									"size": 1
								},
								"rate_limit": {
									"limit": 10,
									"burst": 1
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
								"workerpool": {
									"size": 1
								},
								"rate_limit": {
									"limit": 10,
									"burst": 1
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
								"max_retries": -1,
								"delay": 5,
								"log_flush_timeout_sec": 5,
								"threshold_bytes": 500000,
								"workerpool": {
									"size": 1
								},
								"rate_limit": {
									"limit": 10,
									"burst": 1
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
					"content_type_filter": [],
					"cache": {
						"type": "local",
						"config": {
							"expiration": 2880,
							"cleanup_interval": 300
						}
					}
				}
			]
		},
		{
			"consumer": {
				"name": "mix3_log_api_client",
				"config": {
					"credentials": {
						"auth_disabled": false,
						"token_url": "https://auth.crt.nuance.com/oauth2/token",
						"scope": "log",
						"client_id": "PROVIDE A CLIENT ID",
						"client_secret": "PROVIDE A CLIENT SECRET",
						"auth_timeout_ms": 5000
					},
					"max_retries": -1,
					"retry_delay": 5,
					"record_check_frequency": 6,
					"client_name": "default",
					"consumer_group_id": "01",
					"api_url": "http://localhost:8082",
					"consumer_options": {
						"auto.offset.reset": "earliest"
					}
				}
			},
			"processors": [
				{
					"name": "default",
					"transformer": {
						"type": "default",
						"config": {}
					},
					"storage": [
						{
							"type": "elasticsearch",
							"config": {
								"addresses": [
									"http://localhost:9200"
								],
								"workerpool": {
									"size": 1
								},
								"rate_limit": {
									"limit": 10,
									"burst": 1
								}
							}
						},
						{
							"type": "filesystem",
							"config": {
								"path": "default-processor-logs/customerB-appname-dev",
								"workerpool": {
									"size": 1
								},
								"rate_limit": {
									"limit": 10,
									"burst": 1
								}
							}
						},
						{
							"type": "s3",
							"config": {
								"endpoint": "nuance-bucket.s3.us-east-1.amazonaws.com",
								"access_key_id": "XKIA3SFRSHAXP3MWVOVX",
								"secret_access_key": "X6EFhWAHQQlwnnCKl9MIj6B5Y3RuGyEIarSo3ubx",
								"use_ssl": true,
								"workerpool": {
									"size": 1
								},
								"rate_limit": {
									"limit": 10,
									"burst": 1
								}
							}
						}
					],
					"content_type_filter": [
						"application/x-nuance-nluaas-interpretation.v1+json",
						"application/x-nuance-asr-finalresultresponse.v2+json",
						"application/x-nuance-asr-partialresultresponse.v2+json",
						"application/x-nuance-asr-recognitioninitmessage.v2+json",
						"application/x-nuance-asr-statusresponse.v2+json",
						"application/x-nuance-asr-startspeechresponse.v2+json"
					],
					"cache": {
						"type": "local",
						"config": {
							"expiration": 2880,
							"cleanup_interval": 300
						}
					}
				}
			]
		}
	]`

	yamlData, err := yaml.JSONToYAML([]byte(config))
	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile("config.sample.yaml", yamlData, 0644)
}

func LoadElcConfigYaml() {

	data, err := ioutil.ReadFile("elc.config.yaml")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(data))
	var confArr []interface{}
	err = yml.Unmarshal(data, &confArr)
	if err != nil {
		log.Panic(err)
	}

}
