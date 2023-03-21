package logic

import (
	"context"
	"database/sql"
	"strconv"
	"user/internal/model"

	"rpc-common/user/types/user"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserLogic {
	return &SaveUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserLogic) SaveUser(in *user.UserRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	ctx := context.Background()
	userModel := &model.User{
		Name: sql.NullString{
			String: in.Name,
		},
		Gender: sql.NullString{
			String: in.Gender,
		},
	}
	err := l.svcCtx.UserRepo.Save(ctx, userModel)
	if err != nil {
		return nil, err
	}
	respId := strconv.FormatInt(userModel.Id, 10)
	return &user.UserResponse{
		Id:     respId,
		Name:   in.Name,
		Gender: in.Gender,
	}, nil
}
