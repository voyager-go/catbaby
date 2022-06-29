package model

// RoleCreateInput 角色创建
type RoleCreateInput struct {
	RoleBaseInput
}

// RoleUpdateInput 角色更新
type RoleUpdateInput struct {
	RoleId string `json:"role_id"`
	RoleBaseInput
}

type RoleBaseInput struct {
	Avatar        string `json:"avatar"         ` // 头像
	Name          string `json:"name"           ` // 名称
	Nickname      string `json:"nickname"       ` // 别名
	SurvivalStage string `json:"survival_stage" ` // 生存阶段
	Nationality   string `json:"nationality"    ` // 国籍
	Achievement   string `json:"achievement"    ` // 主要成就
	Gender        uint   `json:"gender"         ` // 是否展示:0女1男
	Content       string `json:"content"        ` // 介绍
	IfShow        uint   `json:"if_show"        ` // 是否展示:0否1是
}
