package bootstrap

import (
	"context"

	"github.com/nferreira/adapter/pkg/adapter"
	"github.com/nferreira/app/pkg/app"
	"github.com/nferreira/app/pkg/service"
	"github.com/nferreira/logging/pkg/logging"
)

type Bootstrap interface {
	WithContext(background context.Context) Bootstrap
	WithLogger(logger logging.Logger) Bootstrap
	WithService(name string, service service.Service) Bootstrap
	WithDependency(name string, dep interface{}) Bootstrap
	WithAdapter(name string, adapter adapter.Adapter) Bootstrap
	ConnectAdapterWithService(adapterId string, rule adapter.BindingRule, service service.BusinessService) Bootstrap
	Boot() (app.App, error)
}

func NewBootstrap(ctx context.Context) Bootstrap {
	return NewDefaultBootstrap(ctx)
}
