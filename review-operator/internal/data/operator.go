package data

import (
	"context"
	pb "review-operator/api/appeal/v1"
	"review-operator/internal/biz"
	"review-operator/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type operatorRepo struct {
	data *Data
	log  *log.Helper
}

func NewOperatorRepo(data *Data, logger log.Logger) biz.OperatorRepo {
	return &operatorRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *operatorRepo) ResolveAppeal(ctx context.Context, resolution *model.Resolution) (string, string, error) {
	res, err := r.data.appealClient.ResolveAppeal(ctx, &pb.ResolveAppealRequest{
		AppealID:   resolution.AppealID,
		OperatorID: resolution.OperatorID,
		Reason:     resolution.Reason,
		Content:    resolution.Content,
		Pictures:   resolution.Pictures,
		Videos:     resolution.Videos,
		Status:     resolution.Status,
	})
	return res.AppealID, res.Status, err
}
