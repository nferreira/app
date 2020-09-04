package app

import (
	"github.com/nferreira/adapter/pkg/adapter"
	"github.com/nferreira/app/pkg/service"
)

type App interface {
	AddAdapter(name string, adapter adapter.Adapter)
	AddService(name string, service service.Service)
	GetService(name string) service.Service
	AddDependency(name string, dep interface{})
	GetDependency(name string) interface{}
	WaitForShutdown()
}