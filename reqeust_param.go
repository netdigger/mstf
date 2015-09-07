package main

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

type ReqParam struct {
	Param
	CaseName  string `json:"case_name"`
	ParamName string `json:"param_name"`
}

type ReqParams map[string]ReqParam
type ParamKey string
type Container map[string]interface{}

func (p ReqParams) GenerateJson() ([]byte, error) {
	container := make(Container)
	for k, v := range p {
		container.SetValue(k, v.GetValue())
	}

	return json.Marshal(&container)
}

func (p Container) SetValue(name string, value interface{}) error {
	list := strings.Split(name, ".")
	if len(list) == 1 {
		p[name] = value
		return nil
	}

	//TODO: mutli level json define.
	/*
		sub := p
		for _, subName := range list {
			item, has := sub[subName]
			if !has {
				item = make(Container)
				sub[subName] = item
				sub := item
				continue
			}
		}
	*/
	return errors.New("unsuport params name")
}

func (p *ReqParam) GetValue() interface{} {
	switch p.Type {
	case "string":
		return p.Value
	case "int":
		return reflect.ValueOf(p.Value).Int()
	}

	return p.Value
}
