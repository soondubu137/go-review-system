package service

import (
	replypb "review-service/api/reply/v1"
	reviewpb "review-service/api/review/v1"
	"review-service/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewReviewService)

type ReviewService struct {
	reviewpb.UnimplementedReviewServer
	replypb.UnimplementedReplyServer

	reviewUC *biz.ReviewerUsecase
	replyUC  *biz.ReplierUsecase
}

func NewReviewService(reviewUC *biz.ReviewerUsecase, replyUC *biz.ReplierUsecase) *ReviewService {
	return &ReviewService{
		reviewUC: reviewUC,
		replyUC:  replyUC,
	}
}
