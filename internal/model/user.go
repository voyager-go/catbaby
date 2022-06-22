package model

// UserRegisterInput 注册参数输入
type UserRegisterInput struct {
	Nickname string // 用户昵称
	Phone    string // 手机号
	Password string // 密码
	Status   int    // 状态:0不正常1正常
}

// UserLoginInput 登录参数输入
type UserLoginInput struct {
	Phone    string // 手机号
	Password string // 密码
}
