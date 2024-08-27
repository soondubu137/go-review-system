package service

import (
	"context"
	pb "review-operator/api/operator/v1"
	"review-operator/internal/biz"
	"review-operator/internal/data/model"
)

type OperatorService struct {
	pb.UnimplementedOperatorServer

	uc *biz.OperatorUsecase
}

func NewOperatorService(uc *biz.OperatorUsecase) *OperatorService {
	return &OperatorService{uc: uc}
}

func (s *OperatorService) ResolveAppeal(ctx context.Context, req *pb.ResolveAppealRequest) (*pb.ResolveAppealReply, error) {
	appealID, status, err := s.uc.ResolveAppeal(ctx, &model.Resolution{
		AppealID:   req.AppealID,
		OperatorID: req.OperatorID,
		Reason:     req.Reason,
		Content:    req.Content,
		Pictures:   req.Pictures,
		Videos:     req.Videos,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ResolveAppealReply{
		AppealID: appealID,
		Status:   status,
	}, err
}
