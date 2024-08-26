package service

import (
	"context"

	pb "review-business/api/business/v1"
	"review-business/internal/biz"
	"review-business/internal/data/model"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer

	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) CreateReply(ctx context.Context, req *pb.CreateReplyRequest) (*pb.CreateReplyReply, error) {
	replyID, err := s.uc.CreateReply(ctx, &model.Reply{
		ReviewID: req.ReviewID,
		UserID:   req.UserID,
		Content:  req.Content,
		Pictures: req.Pictures,
		Videos:   req.Videos,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateReplyReply{
		ReplyID: replyID,
	}, err
}

func (s *BusinessService) CreateAppeal(ctx context.Context, req *pb.CreateAppealRequest) (*pb.CreateAppealReply, error) {
	appealID, status, err := s.uc.CreateAppeal(ctx, &model.Appeal{
		ReviewID: req.ReviewID,
		UserID:   req.UserID,
		Reason:   req.Reason,
		Content:  req.Content,
		Pictures: req.Pictures,
		Videos:   req.Videos,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppealReply{
		AppealID: appealID,
		Status:   status,
	}, err
}
