package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user-score/internal/model"

	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc-common/user-score/types/score"
	"user-score/internal/svc"
)

type SaveUserScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreLogic {
	return &SaveUserScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserScoreLogic) SaveUserScore(in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	// todo: add your logic here and delete this line
	// 从grpc的上下文中生成一个拦截器
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		// 此处失败不用回滚 codes.Internal表示重试
		return nil, status.Error(codes.Internal, err.Error())
	}

	userScore := &model.UserScore{
		UserId: in.UserId,
		Score:  int(in.Score),
	}

	// 此处其实是又起了一个MySQL连接(也就是*sql.Tx) 用这个新连接去做的事务提交
	err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		// 人为制造一个错误以触发callback
		if userScore.Score == 10 {
			return errors.New("user score can't equal 10")
		}

		return l.svcCtx.UserScoreRepo.SaveUserScore(tx, context.Background(), userScore)
	})

	if err != nil {
		// codes.Aborted表示需要回滚
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &score.UserScoreResponse{
		UserId: userScore.UserId,
		Score:  int32(userScore.Score),
	}, nil
}
