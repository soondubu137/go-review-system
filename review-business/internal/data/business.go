package data

import (
	"context"
	pb "review-business/api/reply/v1"
	"review-business/internal/biz"
	"review-business/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *businessRepo) CreateReply(ctx context.Context, reply *model.Reply) (string, error) {
	res, err := r.data.client.CreateReply(ctx, &pb.CreateReplyRequest{
		ReviewID: reply.ReviewID,
		UserID:   reply.UserID,
		Content:  reply.Content,
		Pictures: reply.Pictures,
		Videos:   reply.Videos,
	})
	if err != nil {
		return "", err
	}
	return res.ReplyID, nil
}
