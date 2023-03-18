package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"user/internal/model"
	"user/internal/svc"
	"user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// todo: add your logic here and delete this line
	account := in.Account
	userOrm := &model.User{
		Account: account,
	}
	err := userOrm.FindByName(l.svcCtx.DB)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("查询DB失败: " + err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("账号不存在")
	}

	if userOrm.Password != calcMd5(in.Password) {
		err = errors.New("密码不正确")
		return nil, err
	}
	fmt.Printf("%#v\n", userOrm)
	resp := &user.LoginResponse{
		Code:    200,
		Message: "请求成功",
		Data: &user.LoginData{
			Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.AHsIKlhDgwkzQOn_MDlnsA2KqKRHcpOea3pZpgLj4pw",
			Role: &user.Role{
				Id:   int64(userOrm.Id),
				Name: "super_admin",
			},
		},
	}
	return resp, nil
}

func calcMd5(in string) string {
	h := md5.Sum([]byte(in))
	return hex.EncodeToString(h[:])
}
