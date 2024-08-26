package data

import (
	"review-business/internal/conf"

	pb "review-business/api/reply/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBusinessRepo)

// Data .
type Data struct {
	client pb.ReplyClient
	log    *log.Helper
}

// NewData .
func NewData(c *conf.Data, client pb.ReplyClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{client: client, log: log.NewHelper(logger)}, cleanup, nil
}
