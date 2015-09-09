package main

import ()

type ReqParam struct {
	Param
	From string `json:"from"`
}

type ReqParams map[string]*ReqParam

func (p *ReqParam) Parse(results Results) error {
	if len(p.From) == 0 {
		return nil
	}

	param, err := results.GetParam(p.From)
	if err == nil {
		p.Param = *param
	}
	return err
}
