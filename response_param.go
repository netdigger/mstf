package main

import (
	"reflect"
)

type ResParam struct {
	Param
	Uncheck bool `json:"uncheck"`
}

func (p *ResParam) Equal(param *ResParam) bool {
	if p.Uncheck {
		return true
	}

	return reflect.DeepEqual(p, param)
}
