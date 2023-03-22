package repo

import (
	"context"
	"user-score/internal/model"
)

type UserScoreRepo interface {
	SaveUserScore(ctx context.Context, user *model.UserScore) error
	FindById(ctx context.Context, id int64) (*model.UserScore, error)
}
