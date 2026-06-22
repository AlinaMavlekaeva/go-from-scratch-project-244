package code

import (
	"encoding/json"
	"os"
)

type Config struct {
	host    string
	timeout int
	proxy   string
	follow  bool
	verbose bool
}

func ParseData(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	conf := Config{}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return Config{}, err
	}
	return conf, nil
}
