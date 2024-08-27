package biz

import (
	"context"
	"review-operator/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type OperatorRepo interface {
	ResolveAppeal(context.Context, *model.Resolution) (string, string, error)
}

type OperatorUsecase struct {
	repo OperatorRepo
	log  *log.Helper
}

func NewOperatorUsecase(repo OperatorRepo, logger log.Logger) *OperatorUsecase {
	return &OperatorUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OperatorUsecase) ResolveAppeal(ctx context.Context, appeal *model.Resolution) (string, string, error) {
	return uc.repo.ResolveAppeal(ctx, appeal)
}
