package parsers

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

type Extension interface {
	// Parse(path string) (map[string]any, error)
	Unmarshal(data []byte, v any) error
}

type JSON struct{}
type YAML struct{}

func (ext JSON) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, &v)
}
func (ext YAML) Unmarshal(data []byte, v any) error {
	return yaml.Unmarshal(data, &v)
}

func Parse(path string, e Extension) (map[string]any, error) {
	var conf map[string]any
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}
	err = e.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
