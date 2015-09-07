package main

import (
	"io"
	"log"
	"strings"
)

type Request struct {
	Header     []ReqParam `json:"headers"`
	Parameters []ReqParam `json:"params"`
	Body       ReqParams  `json:"body"`
}

func (p *Request) GetBody(results Results) (io.Reader, error) {
	bodys, err := p.Body.GenerateJson()
	if err != nil {
		return nil, err
	}

	log.Println(string(bodys))
	return strings.NewReader(string(bodys)), nil
}
