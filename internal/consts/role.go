package consts

type RoleGenderState int

const (
	RoleIfShowOk                       = 1 // 展示用户
	RoleIfShowDisabled                 = 0 // 不展示用户
	RoleGenderMan      RoleGenderState = 1 // 男性
	RoleGenderFemale   RoleGenderState = 0 // 女性
	RoleNotFoundErrMsg                 = "未查询到故事角色"
)

func (s RoleGenderState) String() string {
	switch s {
	case RoleGenderMan:
		return "男"
	case RoleGenderFemale:
		return "女"
	default:
		return "Unknown"
	}
}
