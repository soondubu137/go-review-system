package data

import (
	"context"
	"review-service/internal/biz"
	"review-service/internal/data/model"
	"review-service/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
)

type replierRepo struct {
	data *Data
	log  *log.Helper
}

func NewReplierRepo(data *Data, logger log.Logger) biz.ReplierRepo {
	return &replierRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *replierRepo) CreateReply(ctx context.Context, reply *model.Reply) (*model.Reply, error) {
	err := r.data.q.Transaction(func(tx *query.Query) error {
		if err := tx.Reply.WithContext(ctx).Create(reply); err != nil {
			return err
		}
		if _, err := tx.Review.WithContext(ctx).Where(tx.Review.ReviewID.Eq(reply.ReviewID)).Update(tx.Review.HasReply, true); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.log.Errorf("CreateReply error: %v", err)
	}
	return reply, err
}
