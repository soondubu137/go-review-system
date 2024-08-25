package biz

import (
	"context"
	errpb "review-service/api/error/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	"github.com/go-kratos/kratos/v2/log"
)

type ReviewerRepo interface {
	CreateReview(context.Context, *model.Review) (*model.Review, error)
	FindByID(context.Context, int64) (*model.Review, error)
	GetReviewByOrderID(context.Context, int64) (*model.Review, error)
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
	// 3. Look up the corresponding order/product information. (not implemented)
	// 4. Assemble the review data. (not needed)
	// 5. Save the review data to the database.

	r, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil && err.Error() != "record not found" {
		return nil, errpb.ErrorInternal("Internal Server Error")
	}
	if r != nil {
		return nil, errpb.ErrorOrderAlreadyReviewed("Order %d already reviewed.", r.OrderID)
	}

	review.ReviewID = snowflake.GenID()

	return uc.repo.CreateReview(ctx, review)
}

func (uc *ReviewerUsecase) FindByID(ctx context.Context, reviewID int64) (*model.Review, error) {
	uc.log.WithContext(ctx).Debugf("[BIZ] FindByID - req: %v", reviewID)
	return uc.repo.FindByID(ctx, reviewID)
}
