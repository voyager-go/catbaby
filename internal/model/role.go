package model

// RoleCreateInput 故事角色创建
type RoleCreateInput struct {
	RoleBaseInput
}

// RoleUpdateInput 故事角色更新
type RoleUpdateInput struct {
	Id uint
	RoleBaseInput
}

// RoleListInput 故事角色列表
type RoleListInput struct {
	Page   int    // 分页号码
	Size   int    // 分页数量，最大50
	Search string // 搜索关键词
}

type RoleListOutput struct {
	Page  int            `json:"page" description:"页码"`
	Size  int            `json:"size" description:"每页条数"`
	Total int            `json:"total" description:"总条目"`
	List  []RoleBaseItem `json:"list" description:"列表"`
}

type RoleBaseInput struct {
	Avatar        string // 头像
	Name          string // 名称
	Nickname      string // 别名
	SurvivalStage string // 生存阶段
	Nationality   string // 国籍
	Achievement   string // 主要成就
	Gender        uint   // 是否展示:0女1男
	Content       string // 介绍
	IfShow        uint   // 是否展示:0否1是
}

type RoleBaseItem struct {
	Id            uint   `json:"id"`             // 编号
	Avatar        string `json:"avatar"`         // 头像
	Name          string `json:"name"`           // 名称
	Nickname      string `json:"nickname"`       // 别名
	SurvivalStage string `json:"survival_stage"` // 生存阶段
	Nationality   string `json:"nationality"`    // 国籍
	Achievement   string `json:"achievement"`    // 主要成就
	Gender        uint   `json:"gender"`         // 是否展示:0女1男
	GenderText    string `json:"gender_text"`    // 性别 枚举值
	Content       string `json:"content"`        // 介绍
	IfShow        uint   `json:"if_show"`        // 是否展示:0否1是
	IFShowText    string `json:"if_show_text"`   // 是否展示 枚举值
}
