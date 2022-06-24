package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type UserRegisterReq struct {
	g.Meta   `path:"/register" tags:"用户" method:"post" summary:"用户注册"`
	Nickname string `json:"nickname" v:"required#请输入昵称|length:2,16#昵称格式为2-16位" dc:"用户昵称"`
	Password string `json:"password" v:"required#请输入密码|password#密码格式为6-18位" dc:"用户密码"`
	Phone    string `json:"phone"    v:"phone#请输入正确的手机号" dc:"手机号"`
	Status   uint   `json:"status"   v:"in:0,1#请输入符合格式的状态" dc:"状态:0不正常1正常"`
}

type UserRegisterRes struct {
	g.Meta `mime:"text/json" example:"string"`
}

type UserLoginReq struct {
	g.Meta   `path:"/login" tags:"用户" method:"post" summary:"用户登录"`
	Password string `json:"password" v:"required#请输入密码|password#密码格式为6-18位" dc:"密码"`
	Phone    string `json:"phone"    v:"required#请输入手机号|phone#请输入正确的手机号" dc:"手机号"`
}

type UserLoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type UserRefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type UserRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type UserLogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type UserLogoutRes struct{}
