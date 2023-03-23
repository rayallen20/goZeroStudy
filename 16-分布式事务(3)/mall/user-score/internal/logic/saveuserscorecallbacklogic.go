package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc-common/user-score/types/score"
	"user-score/internal/svc"
)

type SaveUserScoreCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreCallbackLogic {
	return &SaveUserScoreCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SaveUserScoreCallback 分布式事务中提交失败之后 回滚时需要执行一个callback
func (l *SaveUserScoreCallbackLogic) SaveUserScoreCallback(in *score.UserScoreRequest) (*score.UserScoreResponse, error) {
	ctx := context.Background()
	err := l.svcCtx.UserScoreRepo.DeleteByUserId(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &score.UserScoreResponse{}, nil
}
