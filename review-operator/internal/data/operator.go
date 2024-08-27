package data

import (
	"context"
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

}
