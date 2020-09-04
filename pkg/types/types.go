package types

import (
	"github.com/nferreira/adapter/pkg/adapter"
	"github.com/nferreira/app/pkg/service"
)

type (
	AdapterRegistry      map[string]adapter.Adapter
	ServiceRegistry      map[string]service.Service
	BindingRuleToService map[adapter.BindingRule]service.BusinessService
	BindingRulesRegistry map[string]BindingRuleToService
	DependencyRegistry   map[string]interface{}
)
