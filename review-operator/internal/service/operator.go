package service

import (
	"context"
	pb "review-operator/api/operator/v1"
	"review-operator/internal/biz"
)

type OperatorService struct {
	pb.UnimplementedOperatorServer

	uc *biz.OperatorUsecase
}

func NewOperatorService(uc *biz.OperatorUsecase) *OperatorService {
	return &OperatorService{uc: uc}
}

func (s *OperatorService) ResolveAppeal(ctx context.Context, req *pb.ResolveAppealRequest) (*pb.ResolveAppealReply, error) {

}
