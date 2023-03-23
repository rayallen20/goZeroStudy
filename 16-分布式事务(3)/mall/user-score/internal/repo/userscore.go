package repo

import (
	"context"
	"database/sql"
	"user-score/internal/model"
)

type UserScoreRepo interface {
	SaveUserScore(tx *sql.Tx, ctx context.Context, user *model.UserScore) error
	FindById(ctx context.Context, id int64) (*model.UserScore, error)
	DeleteByUserId(ctx context.Context, userId int64) error
}
