package data

import (
	pb "review-operator/api/appeal/v1"
	"review-operator/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOperatorRepo)

// Data .
type Data struct {
	appealClient pb.AppealClient
	log          *log.Helper
}

// NewData .
func NewData(c *conf.Data, ac pb.AppealClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{appealClient: ac, log: log.NewHelper(logger)}, cleanup, nil
}
