package service

type ResultBuilder struct {
	Result
}

func NewResultBuilder() *ResultBuilder {
	return &ResultBuilder{
		Result{
			Code:     0,
			Headers:  Headers{},
			ResponseType: JsonResponse,
			Response: nil,
			Error:    nil,
		},
	}
}

func (r *ResultBuilder) WithCode(code int) *ResultBuilder {
	r.Code = code
	return r
}

func (r *ResultBuilder) WithHeaders(headers Headers) *ResultBuilder {
	r.Headers = headers
	return r
}

func (r *ResultBuilder) WithResponseType(responseType ResponseType) *ResultBuilder {
	r.ResponseType = responseType
	return r
}

func (r *ResultBuilder) WithResponseMap(response Response) *ResultBuilder {
	r.Response = response
	return r
}

func (r *ResultBuilder) WithResponseObject(response interface{}) *ResultBuilder {
	r.Response = response
	return r
}

func (r *ResultBuilder) WithError(err error) *ResultBuilder {
	r.Error = err
	return r
}

func (r *ResultBuilder) Build() *Result {
	return &r.Result
}
