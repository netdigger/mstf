package main

import (
	"errors"
	"strings"
)

type Result struct {
	Status   bool
	err      error
	Response Response
}

type Results map[string]Result

func (p Results) GetParam(name string) (*Param, error) {
	list := strings.Split(name, ".")
	if len(list) <= 3 {
		return nil, errors.New("error name: " + name)
	}

	rst, has := p[list[0]]
	if !has {
		return nil, errors.New("there isn't test that's name is " +
			list[0] + " in tests")
	}

	if rst.err != nil {
		return nil, errors.New("can't get param form " + list[0] +
			" test which is error")
	}

	param, err := rst.Response.GetParam(list[1:]...)
	if err != nil {
		return nil, err
	}

	return &param.Param, nil
}
