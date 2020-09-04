package bootstrap

import (
	"context"
	"fmt"
	"runtime"

	"github.com/nferreira/adapter/pkg/adapter"
	"github.com/nferreira/app/pkg/app"
	"github.com/nferreira/app/pkg/service"
	"github.com/nferreira/app/pkg/types"
	"github.com/nferreira/logging/pkg/logging"
)

type DefaultBootstrap struct {
	ctx          context.Context
	logger       logging.Logger
	myApp        *app.DefaultApp
	adapters     types.AdapterRegistry
	services     types.ServiceRegistry
	dependencies types.DependencyRegistry
	bindingRules types.BindingRulesRegistry
}

func NewDefaultBootstrap(ctx context.Context) Bootstrap {
	return &DefaultBootstrap{
		ctx:          ctx,
		myApp:        nil,
		adapters:     types.AdapterRegistry{},
		services:     types.ServiceRegistry{},
		dependencies: types.DependencyRegistry{},
		bindingRules: map[string]types.BindingRuleToService{},
	}
}

func (b *DefaultBootstrap) WithContext(ctx context.Context) Bootstrap {
	b.ctx = ctx
	return b
}

func (b *DefaultBootstrap) WithLogger(logger logging.Logger) Bootstrap {
	b.logger = logger
	b.WithService(logging.LogService, logger)
	return b
}

func (b *DefaultBootstrap) WithService(name string, service service.Service) Bootstrap {
	b.services[name] = service
	return b
}

func (b *DefaultBootstrap) WithDependency(name string, dep interface{}) Bootstrap {
	b.dependencies[name] = dep
	return b
}

func (b *DefaultBootstrap) WithAdapter(name string, adapter adapter.Adapter) Bootstrap {
	b.adapters[name] = adapter
	return b
}

func (b *DefaultBootstrap) ConnectAdapterWithService(adapterId string, rule adapter.BindingRule, service service.BusinessService) Bootstrap {
	adapterBindingRules := b.bindingRules[adapterId]
	if adapterBindingRules == nil {
		adapterBindingRules = make(types.BindingRuleToService)
		b.bindingRules[adapterId] = adapterBindingRules
	}
	adapterBindingRules[rule] = service

	// Notice that we need to register the service
	// This avoid the caller to have to invoke WithService before
	// connecting the service with the adapter
	b.WithService(service.Name(), service)

	return b
}

func (b *DefaultBootstrap) Boot() (_ app.App, err error) {
	b.myApp = app.NewDefaultApp(b.logger)
	b.myApp.Context = context.WithValue(b.ctx, "app", b.myApp)

	b.configRuntime()

	err = b.bootAdapters()
	if err != nil {
		return nil, err
	}

	err = b.bootServices()
	if err != nil {
		return nil, err
	}

	return b.myApp, nil
}

func (b *DefaultBootstrap) bootAdapters() error {
	for id, a := range b.adapters {
		bindingRules := b.bindingRules[id]
		if len(bindingRules) > 0 {
			a.BindRules(bindingRules)
		}
		b.startService(a, "Adapter", id)
		b.myApp.AddAdapter(id, a)
	}
	return nil
}

func (b *DefaultBootstrap) bootServices() error {
	for id, s := range b.services {
		b.startService(s, "Service", id)
		b.myApp.AddService(id, s)
	}
	return nil
}

func (b *DefaultBootstrap) configRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
}

func (b *DefaultBootstrap) startService(s service.Service, serviceType string, id string) {
	go func() {
		if err := s.Start(b.myApp.Context); err != nil {
			panic(fmt.Sprintf("Failed to start %s: %s. Reason: %s", serviceType, id, err.Error()))
		}
	}()
}
