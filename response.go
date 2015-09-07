package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Response struct {
	Status  int               `json:"status"`
	Headers map[string]*Param `json:"header"`
	Bodys   map[string]*Param `json:"body"`
}

func (p *Response) Check(httpResp *http.Response) (*Response, error) {
	resp := &Response{Status: httpResp.StatusCode}
	if err := resp.readBody(httpResp); err != nil {
		return nil, err
	}

	log.Println(resp, p)
	err := p.compare(resp)
	log.Println(resp, p)

	return resp, err
}

func (p *Response) compare(resp *Response) error {
	// TODO check status and headers
	for k, v := range p.Bodys {
		respV, has := resp.Bodys[k]
		if !has {
			return errors.New("have't body parameter:" + k)
		}
		respV.Revision(v.Type)
		log.Println(k, *v, *respV)
		if !v.Uncheck && !reflect.DeepEqual(v, respV) {
			return errors.New("body pararmeter:" + k + " aren't equal")
		}
	}
	return nil
}

func (p *Response) readBody(resp *http.Response) error {
	defer resp.Body.Close()

	if p.Bodys == nil {
		p.Bodys = make(map[string]*Param, 0)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	if len(body) == 0 {
		return nil
	}

	log.Println(string(body))
	bodys := make(map[string]interface{})
	if err := json.Unmarshal(body, &bodys); err != nil {
		log.Println(err)
		return err
	}

	return p.parseBodys("", bodys)
}

func (p *Response) parseBodys(extName string, bodys interface{}) error {
	t := reflect.TypeOf(bodys)
	switch t.Kind() {
	case reflect.Map:
		return p.parseMap(extName, bodys)
	case reflect.Slice:
		return p.parseSlice(extName, bodys)
	}
	return errors.New("unsupport data type:" + t.Kind().String())
}

func (p *Response) parseMap(extName string, bodys interface{}) error {
	value := reflect.ValueOf(bodys)
	keys := value.MapKeys()

	for _, key := range keys {
		name := extName
		if len(name) == 0 {
			name = key.Interface().(string)
		} else {
			name = name + "." + key.Interface().(string)
		}
		v := value.MapIndex(key).Elem()
		switch v.Kind() {
		case reflect.Float64, reflect.String:
			p.Bodys[name] = p.generateParam(v)
		default:
			if err := p.parseBodys(name, v.Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Response) parseSlice(extName string, bodys interface{}) error {
	return errors.New("unsupport slice")
}

func (*Response) generateParam(value reflect.Value) *Param {
	return &Param{
		Type:  value.Kind().String(),
		Value: value.Interface(),
	}
}
