package yaml_test

import (
	"go-mastery/pkg/yaml"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	yaml.LoadConfig("./config.yaml")

}
