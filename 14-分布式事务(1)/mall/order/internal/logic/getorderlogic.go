package logic

import (
	"context"
	"order/internal/svc"
	"order/internal/types"
	"user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	userId := req.Id
	getUserRequest := &user.IdRequest{Id: userId}
	// rpc调用user.GetUser
	userResp, err := l.svcCtx.UserRpc.GetUser(context.Background(), getUserRequest)
	if err != nil {
		return nil, err
	}

	resp = &types.OrderReply{
		Id:       req.Id,
		Name:     "hello order name",
		UserName: userResp.Name,
	}

	return resp, nil
}
