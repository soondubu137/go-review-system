package service

import (
	"context"
	"strconv"

	pb "review-service/api/review/v1"
)

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewReply, error) {
	res, err := s.reviewUC.CreateReview(ctx, createReviewReq2Model(req))
	if err != nil {
		return nil, err
	}
	return &pb.CreateReviewReply{
		ReviewID: strconv.FormatInt(res.ReviewID, 10),
		Status:   res.Status,
	}, nil
}
func (s *ReviewService) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewReply, error) {
	return &pb.UpdateReviewReply{}, nil
}
func (s *ReviewService) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewReply, error) {
	return &pb.DeleteReviewReply{}, nil
}
func (s *ReviewService) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.GetReviewReply, error) {
	return &pb.GetReviewReply{}, nil
}
func (s *ReviewService) ListReview(ctx context.Context, req *pb.ListReviewRequest) (*pb.ListReviewReply, error) {
	return &pb.ListReviewReply{}, nil
}
