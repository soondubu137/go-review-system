package biz

import (
	"context"
	errpb "review-service/api/error/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	"github.com/go-kratos/kratos/v2/log"
)

type AppealerRepo interface {
	CreateAppeal(context.Context, *model.Appeal) (*model.Appeal, error)
	FindByReviewID(context.Context, int64) (*model.Appeal, error)
}

type AppealerUsecase struct {
	repo AppealerRepo
	log  *log.Helper
}

func NewAppealerUsecase(repo AppealerRepo, logger log.Logger) *AppealerUsecase {
	return &AppealerUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *AppealerUsecase) CreateAppeal(ctx context.Context, appeal *model.Appeal, review *model.Review) (*model.Appeal, error) {
	uc.log.WithContext(ctx).Debugf("[BIZ] CreateAppeal - req: %v", appeal)

	// To create an appeal, we follow the following steps:
	// 1. Validate the input data.
	// 2. Horizontal authorization check. (Seller A can only appeal reviews of products sold by Seller A.)
	// 3. If an appeal already exists, check its status.
	// 3.1 If the appeal is pending, update the existing appeal.
	// 3.2 If the appeal is resolved (approved or rejected), return an error.
	// 4. If no appeal exists, generate a unique ID for the appeal.
	// 5. Save the appeal to the appeals table.

	if review.SellerID != appeal.SellerID {
		uc.log.Errorf("Unmatched seller ID: %d", appeal.SellerID)
		return nil, errpb.ErrorUnauthorizedAppeal("Unmatched seller ID: %d", appeal.SellerID)
	}

	oldAppeal, err := uc.repo.FindByReviewID(ctx, review.ReviewID)
	if err != nil && err.Error() != "record not found" {
		uc.log.Errorf("FindByReviewID error: %v", err)
		return nil, errpb.ErrorInternal("Internal Server Error")
	}
	if oldAppeal != nil {
		if oldAppeal.Status != "PENDING" {
			uc.log.Errorf("Appeal %d already resolved", oldAppeal.AppealID)
			return nil, errpb.ErrorAppealAlreadyResolved("Appeal %d already resolved", oldAppeal.AppealID)
		}
		// carry over old fields
		appeal.AppealID = oldAppeal.AppealID
		appeal.Status = oldAppeal.Status
	}

	// only generate a new appeal id if the appeal is new
	if appeal.AppealID == 0 {
		appeal.AppealID = snowflake.GenID()
	}

	return uc.repo.CreateAppeal(ctx, appeal)
}
