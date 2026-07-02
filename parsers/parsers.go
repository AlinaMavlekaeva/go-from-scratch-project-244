package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var ParsersByExt = map[string]func(data []byte, v any) error{
	".json": json.Unmarshal,
	".yml":  yaml.Unmarshal,
	".yaml": yaml.Unmarshal,
}

func Parse(path string) (map[string]any, error) {
	var conf map[string]any
	data, err := os.ReadFile(path)
	if err != nil {
		return conf, err
	}
	ext := filepath.Ext(path)
	f, exists := ParsersByExt[ext]
	if !exists {
		return nil, fmt.Errorf("Unknown extension of file %s", path)
	}
	err = f(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
