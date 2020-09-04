package app

import (
	"context"
	"github.com/nferreira/app/pkg/service"
	"github.com/nferreira/app/pkg/types"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nferreira/adapter/pkg/adapter"
	"github.com/nferreira/logging/pkg/logging"
)

type DefaultApp struct {
	logger logging.Logger
	context.Context
	adapters     types.AdapterRegistry
	services     types.ServiceRegistry
	dependencies types.DependencyRegistry
}

func NewDefaultApp(logger logging.Logger) *DefaultApp {
	return &DefaultApp{
		logger:       logger,
		Context:      nil,
		adapters:     types.AdapterRegistry{},
		services:     types.ServiceRegistry{},
		dependencies: types.DependencyRegistry{},
	}
}

func (a *DefaultApp) AddAdapter(name string, adapter adapter.Adapter) {
	a.adapters[name] = adapter
}

func (a *DefaultApp) AddService(name string, service service.Service) {
	a.services[name] = service
}

func (a *DefaultApp) GetService(name string) service.Service {
	return a.services[name]
}

func (a *DefaultApp) AddDependency(name string, dep interface{}) {
	a.dependencies[name] = dep
}

func (a *DefaultApp) GetDependency(name string) interface{} {
	return a.dependencies[name]
}

func (a *DefaultApp) WaitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Shutdown()
}

func (a *DefaultApp) Shutdown() {
	for _, ad := range a.adapters {
		a.stop(ad)
	}

	for _, s := range a.services {
		a.stop(s)
	}
}

func (a *DefaultApp) stop(s service.Service) {
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(a.Context, timeout)

	quit := make(chan bool, 1)
	go func() {
		_ = s.Stop(ctx)
		quit <- true
	}()

	select {
	case <-time.After(timeout):
		cancel()
	case <-ctx.Done():
		return
	case <-quit:
		return
	}
}
