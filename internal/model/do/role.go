// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta        `orm:"table:role, do:true"`
	Id            interface{} // 主键编号
	Avatar        interface{} // 头像
	Name          interface{} // 名称
	Nickname      interface{} // 别名
	SurvivalStage interface{} // 生存阶段
	Nationality   interface{} // 国籍
	Achievement   interface{} // 主要成就
	Gender        interface{} // 是否展示:0女1男
	Content       interface{} // 介绍
	IfShow        interface{} // 是否展示:0否1是
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
