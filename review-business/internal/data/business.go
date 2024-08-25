package data

import (
	"context"
	"review-business/internal/biz"
	"review-business/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *businessRepo) CreateReply(ctx context.Context, reply *model.Reply) (*model.Reply, error) {
	return nil, nil
}
