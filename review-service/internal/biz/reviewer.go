package biz

import (
	"context"
	"review-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type ReviewerRepo interface {
	CreateReview(context.Context, *model.Review) (*model.Review, error)
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

	// To create a review, we follow the following steps:
	// 1. Validate the input data.
	// 2. Generate a unique ID for the review.
	// 3. Look up the corresponding order/product information.
	// 4. Assemble the review data.
	// 5. Save the review data to the database.

	return uc.repo.CreateReview(ctx, review)
}
