package dao

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user-score/database"
	"user-score/internal/model"
)

// cacheUserIdPrefix 使用缓存时的key前缀
var cacheUserScoreIdPrefix = "cache:user-score:id:"

type UserScoreDao struct {
	Conn *database.DBConn
}

func NewUserScoreDao(conn *database.DBConn) *UserScoreDao {
	return &UserScoreDao{
		Conn: conn,
	}
}

func (u *UserScoreDao) SaveUserScore(ctx context.Context, user *model.UserScore) error {
	sql := fmt.Sprintf("INSERT INTO %s (`user_id`, `score`) VALUES (?, ?)", user.TableName())
	result, err := u.Conn.Conn.ExecCtx(ctx, sql, user.UserId, user.Score)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = id
	return nil
}

func (u *UserScoreDao) FindById(ctx context.Context, id int64) (user *model.UserScore, err error) {
	user = &model.UserScore{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE `id` = ?", user.TableName())
	// 数据在redis中的key名
	userIdKey := fmt.Sprintf("%s:%d", cacheUserScoreIdPrefix, id)
	// 使用带有cache的连接
	err = u.Conn.ConnCache.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, sql, id)
	})

	return user, err
}
