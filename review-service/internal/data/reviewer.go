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

// NewGreeterRepo .
func NewReviewerRepo(data *Data, logger log.Logger) biz.ReviewerRepo {
	return &reviewerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewerRepo) Save(ctx context.Context, g *model.Review) (*model.Review, error) {
	return g, nil
}

func (r *reviewerRepo) Update(ctx context.Context, g *model.Review) (*model.Review, error) {
	return g, nil
}

func (r *reviewerRepo) FindByID(context.Context, int64) (*model.Review, error) {
	return nil, nil
}

func (r *reviewerRepo) ListByHello(context.Context, string) ([]*model.Review, error) {
	return nil, nil
}

func (r *reviewerRepo) ListAll(context.Context) ([]*model.Review, error) {
	return nil, nil
}
