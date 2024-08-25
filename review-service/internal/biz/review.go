package biz

import (
	"context"
	pb "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	"github.com/go-kratos/kratos/v2/log"
)

type ReviewerRepo interface {
	CreateReview(context.Context, *model.Review) (*model.Review, error)
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
	if err != nil {
		uc.log.Errorf("GetReviewByOrderID error: %v", err)
		return nil, pb.ErrorInternal("Internal Server Error")
	}
	if r != nil {
		uc.log.Errorf("Review already exists: %d", r.OrderID)
		return nil, pb.ErrorOrderAlreadyReviewed("Order %d already reviewed.", r.OrderID)
	}

	review.ReviewID = snowflake.GenID()

	return uc.repo.CreateReview(ctx, review)
}
