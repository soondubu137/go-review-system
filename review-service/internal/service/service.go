package service

import (
	appealpb "review-service/api/appeal/v1"
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
	appealpb.UnimplementedAppealServer

	reviewUC *biz.ReviewerUsecase
	replyUC  *biz.ReplierUsecase
	appealUC *biz.AppealerUsecase
}

func NewReviewService(reviewUC *biz.ReviewerUsecase, replyUC *biz.ReplierUsecase, appealUC *biz.AppealerUsecase) *ReviewService {
	return &ReviewService{
		reviewUC: reviewUC,
		replyUC:  replyUC,
		appealUC: appealUC,
	}
}
