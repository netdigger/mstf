package main

import (
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}

func main() {
	log.Println("hello")
	config, err := ReadConfig("./config.json")
	if err != nil {
		return
	}

	dup := make(map[int]string)
	results := make(Results)
	client := CreateClient()

	tests := config.GetTests()
	for index, test := range tests {
		if _, has := results[test.Name]; has {
			dup[index] = test.Name
			continue
		}
		test.Run(client, results)
	}
}

func CreateClient() *http.Client {
	return &http.Client{}
}
