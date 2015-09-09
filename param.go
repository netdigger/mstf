package main

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Param struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func (p *Param) Revision(t string) {
	if strings.HasPrefix(p.Type, "int") {
		p.Type = "int"
		return
	}
	if p.Type == reflect.Float64.String() && t == "int" {
		p.Type = "int"
		return
	}
}

func (p *Param) String() (string, error) {
	v := reflect.ValueOf(p.Value)
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.String && p.Type == "string" {
		return v.String(), nil
	}

	if v.Kind() == reflect.Float64 && p.Type == "float" {
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	}

	if v.Kind() == reflect.Float64 && p.Type == "int" {
		return strconv.FormatInt(int64(v.Float()), 10), nil
	}

	return "", errors.New("type " + p.Type + " don't fit into " +
		v.Kind().String())
}
