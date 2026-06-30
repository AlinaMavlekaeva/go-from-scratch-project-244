package parsers

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseJSON(path string) (map[string]any, error) {
	var conf map[string]any
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
func ParseYAML(path string) (map[string]any, error) {
	var conf map[string]any
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
