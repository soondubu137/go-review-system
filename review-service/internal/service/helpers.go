package service

import (
	"encoding/json"
	pb "review-service/api/review/v1"
	"review-service/internal/data/model"
	"strconv"
)

func parseInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func marshalStrSlice(s []string) string {
	res, _ := json.Marshal(s)
	return string(res)
}

func createReviewReq2Model(req *pb.CreateReviewRequest) *model.Review {
	review := model.Review{
		BuyerID:        parseInt64(req.UserID),
		OrderID:        parseInt64(req.OrderID),
		Rating:         req.Rating,
		ServiceRating:  req.ServiceRating,
		DeliveryRating: req.DeliveryRating,
		Content:        req.Content,
		IsAnonymous:    req.IsAnonymous,
	}
	if len(req.Pictures) > 0 {
		review.Pictures = marshalStrSlice(req.Pictures)
		review.HasMedia = true
	}
	if len(req.Videos) > 0 {
		review.Videos = marshalStrSlice(req.Videos)
		review.HasMedia = true
	}
	return &review
}
