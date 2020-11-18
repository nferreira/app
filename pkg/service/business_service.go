package service

import "context"

type (
	Headers  map[string]interface{}
	Response interface {}
	Params   map[string]interface{}
)

type ResponseType int

const (
	JsonResponse ResponseType = iota
	HtmlResponse
)

type Result struct {
	Code     int
	Headers Headers
	ResponseType ResponseType
	Response Response
	Error    error
}

type BusinessService interface {
	Service
	Name() string
	CreateRequest() interface{}
	Execute(ctx context.Context, params *Params) *Result
}


