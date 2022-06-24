package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/bcrypt"
	"threebody/internal/consts"
	"threebody/internal/dao"
	"threebody/internal/model"
)

type sUser struct{}

func User() *sUser {
	return &sUser{}
}

// Register 注册用户
func (*sUser) Register(ctx context.Context, input model.UserRegisterInput) error {
	user, err := dao.User.FindOneByPhone(ctx, input.Phone)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New(consts.UserPhoneExistErrMsg)
	}
	cryptStrByte, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(cryptStrByte)
	return dao.User.Create(ctx, input)
}

// FindByPhoneAndPassword 根据用户手机号与密码匹配用户
func (*sUser) FindByPhoneAndPassword(ctx context.Context, input model.UserLoginInput) (map[string]interface{}, error) {
	user, err := dao.User.FindOneByPhone(ctx, input.Phone)
	if err != nil || user == nil {
		return nil, errors.New(consts.UserPhoneNotFoundErrMsg)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New(consts.UserPasswordNotMatchErrMsg)
	}
	if user.Status != consts.UserStatusOk {
		return nil, errors.New(consts.UserStatusDisabledErrMsg)
	}
	return g.Map{
		"id":       user.Id,
		"nickname": user.Nickname,
		"phone":    user.Phone,
	}, nil
}
