package main

type Result struct {
	Status   bool
	err      error
	Response Response
}

type Results map[string]Result
