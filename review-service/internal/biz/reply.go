package biz

import (
	"context"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	errpb "review-service/api/error/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type ReplierRepo interface {
	CreateReply(context.Context, *model.Reply) (*model.Reply, error)
}

type ReplierUsecase struct {
	repo ReplierRepo
	log  *log.Helper
}

func NewReplierUsecase(repo ReplierRepo, logger log.Logger) *ReplierUsecase {
	return &ReplierUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ReplierUsecase) CreateReply(ctx context.Context, reply *model.Reply, review *model.Review) (*model.Reply, error) {
	uc.log.WithContext(ctx).Debugf("[BIZ] CreateReply - req: %v", reply)

	// To create a reply, we follow the following steps:
	// 1. Validate the input data.
	// 2. Horizontal authorization check. (Seller A can only reply to reviews of products sold by Seller A.)
	// 3. Generate a unique ID for the reply.
	// 4. Update replies table and reviews table (transaction).

	if review.HasReply {
		uc.log.Errorf("Review %d already has a reply", review.ReviewID)
		return nil, errpb.ErrorReviewAlreadyReplied("Review %d already has a reply.", review.ReviewID)
	}

	if review.SellerID != reply.SellerID {
		uc.log.Errorf("Unmatched seller ID: %d", reply.SellerID)
		return nil, errpb.ErrorUnauthorizedReply("Unmatched seller ID: %d", reply.SellerID)
	}

	reply.ReplyID = snowflake.GenID()

	return uc.repo.CreateReply(ctx, reply)
}
