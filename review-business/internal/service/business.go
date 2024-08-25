package service

import (
	"context"

	pb "review-business/api/business/v1"
	"review-business/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer

	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) CreateReply(ctx context.Context, req *pb.CreateReplyRequest) (*pb.CreateReplyReply, error) {
	return &pb.CreateReplyReply{}, nil
}
