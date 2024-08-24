package biz

import (
	"context"
	"review-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type ReviewerRepo interface {
	Save(context.Context, *model.Review) (*model.Review, error)
}

type ReviewerUsecase struct {
	repo ReviewerRepo
	log  *log.Helper
}

func NewReviewerUsecase(repo ReviewerRepo, logger log.Logger) *ReviewerUsecase {
	return &ReviewerUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ReviewerUsecase) CreateReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	uc.log.WithContext(ctx).Debugf("[BIZ] CreateReview - req: %v", review)
	return uc.repo.Save(ctx, review)
}
