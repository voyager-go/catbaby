package role

import (
	"context"
	"threebody/internal/dao"
	"threebody/internal/model"
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
