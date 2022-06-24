package controller

import (
	"context"
	"threebody/api/v1"
	"threebody/internal/model"
	"threebody/internal/service"
)

type user struct{}

var User = user{}

// Register 用户注册
func (*user) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	err = service.User().Register(ctx, model.UserRegisterInput{
		Nickname: req.Nickname,
		Password: req.Password,
		Phone:    req.Phone,
		Status:   req.Status,
	})
	return
}

// Login 发放令牌
func (*user) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	res = &v1.UserLoginRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}

// RefreshToken 刷新令牌
func (*user) RefreshToken(ctx context.Context, req *v1.UserRefreshTokenReq) (res *v1.UserRefreshTokenRes, err error) {
	res = &v1.UserRefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

// Logout 退出
func (*user) Logout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
