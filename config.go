package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Url   string `json:"url"`
	Tests []Test `json:"tests"`
}

func ReadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Println("json error:", err)
		return nil, err
	}

	for index, _ := range config.Tests {
		config.Tests[index].Url = config.Url + config.Tests[index].Url
	}

	return &config, nil
}

func (p *Config) GetTests() []Test {
	return p.Tests
}
