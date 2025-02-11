package service

import (
	"encoding/json"
	appealpb "review-service/api/appeal/v1"
	replypb "review-service/api/reply/v1"
	reviewpb "review-service/api/review/v1"
	"review-service/internal/data/model"
	"strconv"
	"time"
)

func parseInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func marshalStrSlice(s []string) string {
	res, _ := json.Marshal(s)
	return string(res)
}

func createReviewReq2Model(req *reviewpb.CreateReviewRequest) *model.Review {
	review := model.Review{
		BuyerID:        parseInt64(req.UserID),
		OrderID:        parseInt64(req.OrderID),
		Rating:         req.Rating,
		ServiceRating:  req.ServiceRating,
		DeliveryRating: req.DeliveryRating,
		Content:        req.Content,
		IsAnonymous:    req.IsAnonymous,
		Pictures:       "[]",
		Videos:         "[]",
		Tags:           "[]",
		Snapshot:       "{}",
		ExtJSON:        "{}",
		CtrlJSON:       "{}",
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

func createReplyReq2Model(req *replypb.CreateReplyRequest) *model.Reply {
	reply := model.Reply{
		ReviewID: parseInt64(req.ReviewID),
		SellerID: parseInt64(req.UserID),
		Content:  req.Content,
		Pictures: "[]",
		Videos:   "[]",
		ExtJSON:  "{}",
		CtrlJSON: "{}",
	}
	if len(req.Pictures) > 0 {
		reply.Pictures = marshalStrSlice(req.Pictures)
	}
	if len(req.Videos) > 0 {
		reply.Videos = marshalStrSlice(req.Videos)
	}

	return &reply
}

func createAppealReq2Model(req *appealpb.CreateAppealRequest) *model.Appeal {
	appeal := model.Appeal{
		ReviewID: parseInt64(req.ReviewID),
		SellerID: parseInt64(req.UserID),
		Content:  req.Content,
		Pictures: "[]",
		Videos:   "[]",
		ExtJSON:  "{}",
		CtrlJSON: "{}",
	}
	if len(req.Pictures) > 0 {
		appeal.Pictures = marshalStrSlice(req.Pictures)
	}
	if len(req.Videos) > 0 {
		appeal.Videos = marshalStrSlice(req.Videos)
	}

	return &appeal
}

func updateAppealWithResolution(req *appealpb.ResolveAppealRequest, appeal *model.Appeal) *model.Appeal {
	operatorID := parseInt64(req.OperatorID)
	currTime := time.Now()

	appeal.OpProcessBy = &operatorID
	appeal.OpProcessAt = &currTime
	appeal.Reason = req.Reason
	appeal.Content = req.Content
	appeal.Status = req.Status
	if len(req.Pictures) > 0 {
		appeal.Pictures = marshalStrSlice(req.Pictures)
	}
	if len(req.Videos) > 0 {
		appeal.Videos = marshalStrSlice(req.Videos)
	}

	return appeal
}
