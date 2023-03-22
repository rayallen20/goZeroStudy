package logic

import (
	"context"
	"user-score/internal/model"

	"rpc-common/user-score/types/score"
	"user-score/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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
	userScore := &model.UserScore{
		UserId: in.UserId,
		Score:  int(in.Score),
	}
	err := l.svcCtx.UserScoreRepo.SaveUserScore(context.Background(), userScore)
	if err != nil {
		return nil, err
	}
	return &score.UserScoreResponse{
		UserId: userScore.UserId,
		Score:  int32(userScore.Score),
	}, nil
}
