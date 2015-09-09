package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Url   string  `json:"url"`
	Tests []*Test `json:"tests"`
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

	for _, v := range config.Tests {
		v.Url = config.Url + v.Url
	}

	return &config, nil
}

func (p *Config) GetTests() []*Test {
	return p.Tests
}
