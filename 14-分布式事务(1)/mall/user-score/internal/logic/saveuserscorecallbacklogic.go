package logic

import (
	"context"
	"fmt"
	"rpc-common/user-score/types/score"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	fmt.Println("user-score saveUserScore callback...")
	return &score.UserScoreResponse{}, nil
}
