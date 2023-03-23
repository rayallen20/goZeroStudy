package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"user/internal/model"

	"rpc-common/user/types/user"
	"user/internal/svc"

	_ "github.com/dtm-labs/driver-gozero"
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
	// 此处从grpc的上下文中生成一个拦截器
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		// 此处失败不用回滚 codes.Internal表示重试
		return nil, status.Error(codes.Internal, err.Error())
	}

	id, _ := strconv.ParseInt(in.Id, 10, 64)
	ctx := context.Background()
	userModel := &model.User{
		Id: id,
		Name: sql.NullString{
			String: in.Name,
		},
		Gender: sql.NullString{
			String: in.Gender,
		},
	}

	// 此处其实是又起了一个MySQL连接(*sql.Tx) 用这个新连接去做的事务提交
	err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		return l.svcCtx.UserRepo.Save(tx, ctx, userModel)
	})

	if err != nil {
		// codes.Aborted表示需要回滚
		return nil, status.Error(codes.Aborted, err.Error())
	}

	respId := strconv.FormatInt(userModel.Id, 10)
	return &user.UserResponse{
		Id:     respId,
		Name:   in.Name,
		Gender: in.Gender,
	}, nil
}

func (l *SaveUserLogic) SaveUserCallback(in *user.UserRequest) (*user.UserResponse, error) {
	// 在回调中进行补偿操作 如果事务是插入数据 那么回调中就是删除数据
	// 此处以回滚为例 删除数据
	ctx := context.Background()
	id, _ := strconv.ParseInt(in.Id, 10, 64)
	err := l.svcCtx.UserRepo.DeleteById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{}, nil
}
