package data

import (
	"errors"
	"review-service/internal/conf"
	"review-service/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewReviewerRepo, NewReplierRepo)

type Data struct {
	q   *query.Query
	log *log.Helper
}

// Provider of Data.
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	query.SetDefault(db)
	return &Data{q: query.Q, log: log.NewHelper(logger)}, cleanup, nil
}

// NewDB is the provider of DB connection.
func NewDB(cfg *conf.Data) (*gorm.DB, error) {
	if cfg.Database.Driver != "mysql" {
		return nil, errors.New("only mysql supported")
	}
	return gorm.Open(mysql.Open(cfg.Database.Source))
}
