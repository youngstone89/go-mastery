package yaml

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

func DoConvertJsonToYaml() {

	type jsonObj map[string]interface{}

	var jsonData = make(jsonObj)

	jsonData["hi"] = "there"

	data, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	yamlD, err := yaml.JSONToYAML(data)
	if err != nil {
		fmt.Errorf("err %v", err)
	}
	ioutil.WriteFile("ConvertedFromJson.yaml", yamlD, 0644)
}
