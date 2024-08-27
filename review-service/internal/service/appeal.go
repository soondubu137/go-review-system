package service

import (
	"context"
	pb "review-service/api/appeal/v1"
	errpb "review-service/api/error/v1"
	"strconv"
)

func (s *ReviewService) CreateAppeal(ctx context.Context, req *pb.CreateAppealRequest) (*pb.CreateAppealReply, error) {
	// we first need to find the review this appeal is for
	reviewID, _ := strconv.ParseInt(req.ReviewID, 10, 64)
	review, err := s.reviewUC.FindByID(ctx, reviewID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errpb.ErrorInvalidReviewId("Invalid Review ID")
		}
		return nil, errpb.ErrorInternal("Internal Server Error")
	}

	res, err := s.appealUC.CreateAppeal(ctx, createAppealReq2Model(req), review)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAppealReply{
		AppealID: strconv.FormatInt(res.AppealID, 10),
		Status:   res.Status,
	}, nil
}

func (s *ReviewService) ResolveAppeal(ctx context.Context, req *pb.ResolveAppealRequest) (*pb.ResolveAppealReply, error) {
	// we first need to find the appeal to be resolved
	appealID, _ := strconv.ParseInt(req.AppealID, 10, 64)
	appeal, err := s.appealUC.FindByID(ctx, appealID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errpb.ErrorInvalidAppealId("Invalid Appeal ID")
		}
		return nil, errpb.ErrorInternal("Internal Server Error")
	}

	// if the appeal is already resolved, return an error
	if appeal.Status != "PENDING" {
		return nil, errpb.ErrorAppealAlreadyResolved("Appeal %d already resolved", appeal.AppealID)
	}

	res, err := s.appealUC.ResolveAppeal(ctx, updateAppealWithResolution(req, appeal))
	if err != nil {
		return nil, err
	}

	return &pb.ResolveAppealReply{
		AppealID: strconv.FormatInt(res.AppealID, 10),
		Status:   res.Status,
	}, nil
}
