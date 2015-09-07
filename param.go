package main

import (
	"reflect"
	"strings"
)

type Param struct {
	Type    string      `json:"type"`
	Value   interface{} `json:"value"`
	Uncheck bool        `json:"uncheck"`
}

func (p *Param) Revision(t string) {
	if strings.HasPrefix(p.Type, "int") {
		p.Type = "int"
	}
	if p.Type != reflect.Float64.String() || t != "int" {
		return
	}

	p.Type = "int"
	v := reflect.ValueOf(p.Value)
	p.Value = int(v.Float())
}
