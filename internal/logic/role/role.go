package role

import (
	"context"
	"threebody/internal/consts"
	"threebody/internal/dao"
	"threebody/internal/model"
	"threebody/internal/model/entity"
	"threebody/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Create 创建故事角色记录
func (*sRole) Create(ctx context.Context, in model.RoleCreateInput) error {
	_, err := dao.Role.Ctx(ctx).Data(in).Insert()
	return err
}

// Update 更新故事角色记录
func (*sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	_, err := dao.Role.Ctx(ctx).Data(in).Save()
	return err
}

func (*sRole) List(ctx context.Context, in model.RoleListInput) (out *model.RoleListOutput, err error) {
	var (
		m    = dao.Role.Ctx(ctx)
		list []*entity.Role
	)
	out = &model.RoleListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	m = m.Where(dao.Role.Columns().IfShow, consts.RoleIfShowOk)
	listModel := m.Page(in.Page, in.Size)
	if err = listModel.Scan(&list); err != nil {
		return
	}
	if len(list) == 0 {
		return
	}
	out.Total, err = m.Count()
	if err != nil {
		return
	}
	return
}
