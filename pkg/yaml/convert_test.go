package yaml_test

import (
	"go-mastery/pkg/yaml"
	"testing"
)

func TestDoConvertJsonToYaml(t *testing.T) {

	yaml.DoConvertJsonToYaml()

}

func TestUnMarshallYamlToMap(t *testing.T) {

	yaml.UnMarshallYamlToMap()

}

func TestUnMarshalJsonYamlStruct(t *testing.T) {
	yaml.UnMarshalJsonYamlStruct()
}

func TestConvertJsonConfigToYaml(t *testing.T) {
	yaml.ConvertJsonConfigToYaml()

}

func TestLoadElcConfigYaml(t *testing.T) {
	yaml.LoadElcConfigYaml()

}
