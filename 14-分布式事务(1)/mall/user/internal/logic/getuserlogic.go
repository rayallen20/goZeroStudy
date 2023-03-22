package logic

import (
	"context"
	"rpc-common/user/types/user"
	"strconv"
	"user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	id, _ := strconv.Atoi(in.Id)
	logx.Info("get user call\n")
	ctx := context.Background()
	userModel, err := l.svcCtx.UserRepo.FindById(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	resp := &user.UserResponse{
		Id:     strconv.FormatInt(userModel.Id, 10),
		Name:   userModel.Name.String,
		Gender: userModel.Gender.String,
	}
	return resp, nil
}
