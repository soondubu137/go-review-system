package service

import (
	"context"
	errpb "review-service/api/error/v1"
	pb "review-service/api/reply/v1"

	"strconv"
)

func (s *ReviewService) CreateReply(ctx context.Context, req *pb.CreateReplyRequest) (*pb.CreateReplyReply, error) {
	// we first need to find the review this reply is for
	reviewID, _ := strconv.ParseInt(req.ReviewID, 10, 64)
	review, err := s.reviewUC.FindByID(ctx, reviewID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errpb.ErrorInvalidReviewId("Invalid Review ID")
		}
		return nil, errpb.ErrorInternal("Internal Server Error")
	}

	res, err := s.replyUC.CreateReply(ctx, createReplyReq2Model(req), review)
	if err != nil {
		return nil, err
	}

	return &pb.CreateReplyReply{
		ReplyID: strconv.FormatInt(res.ReplyID, 10),
	}, nil
}
