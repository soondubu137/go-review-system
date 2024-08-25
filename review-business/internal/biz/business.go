package biz

import (
	"context"
	"review-business/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type BusinessRepo interface {
	CreateReply(context.Context, *model.Reply) (*model.Reply, error)
}

type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}
