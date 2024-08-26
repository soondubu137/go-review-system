package data

import (
	"review-business/internal/conf"

	appealpb "review-business/api/appeal/v1"
	replypb "review-business/api/reply/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBusinessRepo)

// Data .
type Data struct {
	appealClient appealpb.AppealClient
	replyClient  replypb.ReplyClient
	log          *log.Helper
}

// NewData .
func NewData(c *conf.Data, ac appealpb.AppealClient, rc replypb.ReplyClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		appealClient: ac,
		replyClient:  rc,
		log:          log.NewHelper(logger),
	}, cleanup, nil
}
