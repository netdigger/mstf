package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

type Test struct {
	Name     string   `json:"name"`
	Url      string   `json:"url"`
	Method   string   `json:"method"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

func (p *Test) Run(client *http.Client, results Results) {
	p.showStartLog()
	defer p.showFinishLog()

	req, err := p.newHttpRequest(results)
	if err != nil {
		p.doError(results, err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		p.doError(results, err)
		return
	}

	var res *Response
	res, err = p.Response.Check(resp)
	if err != nil {
		p.doError(results, err)
		return
	}

	results[p.Name] = Result{Response: *res}
}

func (p *Test) showStartLog() {
	log.Println("start test cast:", p.Url, p.Method)
}

func (p *Test) showFinishLog() {
	log.Println("finish test cast:", p.Url, p.Method)
}

func (p *Test) doTest(client *http.Client, result *Result) error {
	return nil
}

func (p *Test) doError(results Results, err error) {
	log.Println(err)
	results[p.Name] = Result{err: err}
}

func (p *Test) newHttpRequest(results Results) (*http.Request, error) {
	var reader io.Reader
	var err error
	url := p.Url
	switch p.Method {
	case "POST":
		reader, err = p.Request.GetBody(results)
	case "GET":
		url, err = p.Request.GetUrl(p.Url, results)
	case "PUT":
		reader, err = p.Request.GetBody(results)
	default:
		err = errors.New("unsupport method:" + p.Method)
	}
	if err != nil {
		return nil, err
	}

	return http.NewRequest(p.Method, url, reader)
}
