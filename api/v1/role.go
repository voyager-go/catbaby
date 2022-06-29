package v1

import "github.com/gogf/gf/v2/frame/g"

type RoleBaseReq struct {
	Avatar        string `json:"avatar" dc:"头像"`
	Name          string `json:"name" v:"required#请输入名称|length:1,30#昵称格式为1-30位" dc:"名称"`
	Nickname      string `json:"nickname" v:"length:1,30#昵称格式为1-30位" dc:"昵称"`
	SurvivalStage string `json:"survival_stage" dc:"生存阶段"`
	Nationality   string `json:"nationality" dc:"国籍"`
	Achievement   string `json:"achievement" dc:"主要成就"`
	Gender        uint   `json:"gender" v:"in:0,1#请选择性别" dc:"状态:0女1男"`
	Content       string `json:"content" v:"required#请输入人物介绍" dc:"人物介绍"`
	IfShow        uint   `json:"if_show" v:"in:0,1#请选择是否展示" dc:"状态:0否1是"`
}

type RoleCreateReq struct {
	g.Meta `path:"/backend/role/create" tags:"三体" method:"post" summary:"故事角色"`
	RoleBaseReq
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" tags:"三体" method:"put" summary:"故事角色"`
	RoleId uint `json:"role_id" v:"min:1#请选择需要修改的角色"`
	RoleBaseReq
}

type RoleRes struct {
	g.Meta `mime:"text/json" example:"string"`
}
