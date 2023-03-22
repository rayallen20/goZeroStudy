package dao

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/database"
	"user/internal/model"
)

// cacheUserIdPrefix 使用缓存时的key前缀
var cacheUserIdPrefix = "cache:user:id:"

type UserDao struct {
	Conn *database.DBConn
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		Conn: conn,
	}
}

func (u *UserDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("INSERT INTO %s (`name`, `gender`) VALUES (?, ?)", user.TableName())
	result, err := u.Conn.Conn.ExecCtx(ctx, sql, user.Name.String, user.Gender.String)
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

func (u *UserDao) FindById(ctx context.Context, id int64) (user *model.User, err error) {
	user = &model.User{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE `id` = ?", user.TableName())
	// 数据在redis中的key名
	userIdKey := fmt.Sprintf("%s:%d", cacheUserIdPrefix, id)
	// 使用带有cache的连接
	err = u.Conn.ConnCache.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, sql, id)
	})

	return user, err
}
