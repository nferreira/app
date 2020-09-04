package service

const ExecutionContextKey = "ExecutionContext"

type ExecutionContext struct {
	CorrelationId string
	App interface{}
}

func NewExecutionContext(correlationId string, app interface{}) *ExecutionContext {
	return &ExecutionContext{
		CorrelationId: correlationId,
		App: app,
	}
}