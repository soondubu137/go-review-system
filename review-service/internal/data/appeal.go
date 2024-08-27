package data

import (
	"context"
	"review-service/internal/biz"
	"review-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type appealerRepo struct {
	data *Data
	log  *log.Helper
}

func NewAppealerRepo(data *Data, logger log.Logger) biz.AppealerRepo {
	return &appealerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *appealerRepo) CreateAppeal(ctx context.Context, appeal *model.Appeal) (*model.Appeal, error) {
	err := r.data.q.Appeal.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "review_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"status":   appeal.Status,
				"reason":   appeal.Reason,
				"content":  appeal.Content,
				"pictures": appeal.Pictures,
				"videos":   appeal.Videos,
			}),
		}).
		Create(appeal)
	return appeal, err
}

func (r *appealerRepo) FindByID(ctx context.Context, appealID int64) (*model.Appeal, error) {
	Appeal := r.data.q.Appeal
	res, err := Appeal.WithContext(ctx).Where(Appeal.AppealID.Eq(appealID)).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *appealerRepo) FindByReviewID(ctx context.Context, reviewID int64) (*model.Appeal, error) {
	Appeal := r.data.q.Appeal
	res, err := Appeal.WithContext(ctx).Where(Appeal.ReviewID.Eq(reviewID)).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *appealerRepo) ResolveAppeal(ctx context.Context, appeal *model.Appeal) (*model.Appeal, error) {
	Appeal := r.data.q.Appeal
	_, err := Appeal.WithContext(ctx).Where(Appeal.AppealID.Eq(appeal.AppealID)).Updates(appeal)
	return appeal, err
}
