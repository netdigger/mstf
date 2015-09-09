package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"strings"
)

type Request struct {
	Header     ReqParams `json:"headers"`
	Parameters ReqParams `json:"params"`
	Body       ReqParams `json:"body"`
}

// Generate the full request url from the base url and parameters in body.
func (p *Request) GetUrl(base string, results Results) (string, error) {
	var url string
	for k, v := range p.Body {
		if err := v.Parse(results); err != nil {
			return "", err
		}

		if value, err := v.String(); err == nil {
			url = url + "&" + k + "=" + value
		} else {
			return "", err
		}
	}
	log.Println(url)

	return base + "?" + strings.TrimPrefix(url, "&"), nil
}

func (p *Request) GetBody(results Results) (io.Reader, error) {
	data := make(map[string]interface{})
	for k, v := range p.Body {
		if err := v.Parse(results); err != nil {
			return nil, err
		}

		data[k] = v.Value
	}
	bodys, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	log.Println(string(bodys))
	return bytes.NewReader(bodys), nil
}
