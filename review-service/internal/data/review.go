package data

import (
	"context"

	"review-service/internal/biz"
	"review-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewerRepo struct {
	data *Data
	log  *log.Helper
}

func NewReviewerRepo(data *Data, logger log.Logger) biz.ReviewerRepo {
	return &reviewerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewerRepo) CreateReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	err := r.data.q.Review.WithContext(ctx).Create(review)
	if err != nil {
		r.log.Errorf("CreateReview error: %v", err)
	}
	return review, err
}

func (r *reviewerRepo) GetReviewByOrderID(ctx context.Context, orderID int64) (*model.Review, error) {
	R := r.data.q.Review
	review, err := R.WithContext(ctx).Where(R.OrderID.Eq(orderID)).First()
	if err != nil {
		r.log.Errorf("GetReviewByOrderID error: %v", err)
		return nil, err
	}
	return review, nil
}

func (r *reviewerRepo) Update(ctx context.Context, review *model.Review) (*model.Review, error) {
	return review, nil
}

func (r *reviewerRepo) FindByID(ctx context.Context, reviewID int64) (*model.Review, error) {
	R := r.data.q.Review
	review, err := R.WithContext(ctx).Where(R.ReviewID.Eq(reviewID)).First()
	if err != nil {
		r.log.Errorf("FindByID error: %v", err)
		return nil, err
	}
	return review, nil
}

func (r *reviewerRepo) ListAll(context.Context) ([]*model.Review, error) {
	return nil, nil
}
