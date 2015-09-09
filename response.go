package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Response struct {
	Status  int                  `json:"status"`
	Headers map[string]*ResParam `json:"header"`
	Bodys   map[string]*ResParam `json:"body"`
}

func (p *Response) Check(httpResp *http.Response) (*Response, error) {
	resp := &Response{Status: httpResp.StatusCode}
	if err := resp.readBody(httpResp); err != nil {
		return nil, err
	}

	err := p.compare(resp)
	return resp, err
}

func (p *Response) GetParam(names ...string) (*ResParam, error) {
	var params map[string]*ResParam
	switch names[0] {
	case "body":
		params = p.Bodys
	case "header":
		params = p.Headers
	default:
		return nil, errors.New("unkown struct field name of response")
	}

	param, has := params[strings.Join(names[1:], ".")]
	if !has {
		return nil, errors.New("unknown name of response parameters")
	}

	return param, nil
}

func (p *Response) compare(resp *Response) error {
	// TODO check status and headers
	for k, v := range p.Bodys {
		respV, has := resp.Bodys[k]
		if !has {
			return errors.New("have't body parameter:" + k)
		}
		respV.Revision(v.Type)
		if !v.Equal(respV) {
			log.Println(k, *v, *respV)
			return errors.New("body pararmeter:" + k + " aren't equal")
		}
	}
	return nil
}

func (p *Response) readBody(resp *http.Response) error {
	defer resp.Body.Close()

	if p.Bodys == nil {
		p.Bodys = make(map[string]*ResParam, 0)
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

	return p.parseBodys("", reflect.ValueOf(bodys))
}

func (p *Response) parseBodys(extName string, v reflect.Value) error {
	kind := v.Kind()
	if kind == reflect.Interface || kind == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Map:
		return p.parseMap(extName, v)
	case reflect.Slice:
		return p.parseSlice(extName, v)
	}
	return errors.New("unsupport data type:" + v.Kind().String())
}

func (p *Response) parseMap(extName string, value reflect.Value) error {
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
			if err := p.parseBodys(name, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Response) parseSlice(extName string, v reflect.Value) error {
	for i := 0; i < v.Len(); i++ {
		name := extName + "[" + strconv.FormatInt(int64(i), 10) + "]"
		item := v.Index(i)
		switch item.Kind() {
		case reflect.Float64, reflect.String:
			p.Bodys[name] = p.generateParam(v)
		default:
			if err := p.parseBodys(name, item); err != nil {
				log.Println(name, err)
				return err
			}
		}
	}
	return nil
}

func (*Response) generateParam(value reflect.Value) *ResParam {
	return &ResParam{
		Param: Param{
			Type:  value.Kind().String(),
			Value: value.Interface(),
		},
	}
}
