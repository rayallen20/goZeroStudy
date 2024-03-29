// Code generated by goctl. DO NOT EDIT.
// Source: user_score.proto

package server

import (
	"context"
	score2 "rpc-common/user-score/types/score"

	"user-score/internal/logic"
	"user-score/internal/svc"
)

type UserScoreServer struct {
	svcCtx *svc.ServiceContext
	score2.UnimplementedUserScoreServer
}

func NewUserScoreServer(svcCtx *svc.ServiceContext) *UserScoreServer {
	return &UserScoreServer{
		svcCtx: svcCtx,
	}
}

func (s *UserScoreServer) SaveUserScore(ctx context.Context, in *score2.UserScoreRequest) (*score2.UserScoreResponse, error) {
	l := logic.NewSaveUserScoreLogic(ctx, s.svcCtx)
	return l.SaveUserScore(in)
}

// saveUserScoreCallback 分布式事务中提交失败之后 回滚时需要执行一个callback
func (s *UserScoreServer) SaveUserScoreCallback(ctx context.Context, in *score2.UserScoreRequest) (*score2.UserScoreResponse, error) {
	l := logic.NewSaveUserScoreCallbackLogic(ctx, s.svcCtx)
	return l.SaveUserScoreCallback(in)
}
