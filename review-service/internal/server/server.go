package server

import (
	"review-service/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(c *conf.Registry) registry.Registrar {
	cfg := api.DefaultConfig()
	cfg.Address = c.Consul.Addr
	cfg.Scheme = c.Consul.Scheme

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return consul.New(client, consul.WithHealthCheck(true))
}
