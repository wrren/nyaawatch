package main

import (
	"encoding/json"
)

type Series struct {
	Name    string
	Subber  string
	Quality string
	Episode int
}

type WatchConfig struct {
	URL       string
	Directory string
	Refresh   int
	Series    []Series
	Regexes   []string
}

func ReadConfig(in []byte) (WatchConfig, error) {
	var config WatchConfig
	err := json.Unmarshal(in, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
