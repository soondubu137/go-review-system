package service

import (
	"context"
	pb "review-business/api/reply/v1"
	"review-business/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

func NewDiscovery(c *conf.Registry) registry.Discovery {
	cfg := api.DefaultConfig()
	cfg.Address = c.Consul.Addr
	cfg.Scheme = c.Consul.Scheme

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return consul.New(client)
}

func NewReviewServiceClient(disc registry.Discovery, c *conf.Registry) pb.ReplyClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+c.ServiceName),
		grpc.WithDiscovery(disc),
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	)
	if err != nil {
		panic(err)
	}
	return pb.NewReplyClient(conn)
}
