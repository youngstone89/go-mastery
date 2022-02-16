package yaml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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
