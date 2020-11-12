package service

import "context"

type (
	Headers  map[string]interface{}
	Response interface {}
	Params   map[string]interface{}
)

type Result struct {
	Code     int
	Headers  Headers
	Response Response
	Error    error
}

type BusinessService interface {
	Service
	Name() string
	CreateRequest() interface{}
	Execute(ctx context.Context, params *Params) *Result
}


