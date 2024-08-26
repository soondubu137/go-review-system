package biz

import (
	"context"
	"review-business/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type BusinessRepo interface {
	CreateReply(context.Context, *model.Reply) (string, error)
	CreateAppeal(context.Context, *model.Appeal) (string, string, error)
}

type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BusinessUsecase) CreateReply(ctx context.Context, reply *model.Reply) (string, error) {
	return uc.repo.CreateReply(ctx, reply)
}

func (uc *BusinessUsecase) CreateAppeal(ctx context.Context, appeal *model.Appeal) (string, string, error) {
	return uc.repo.CreateAppeal(ctx, appeal)
}
