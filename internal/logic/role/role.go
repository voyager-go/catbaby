package role

import (
	"context"
	"errors"
	"threebody/internal/consts"
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

// List 故事角色列表
func (*sRole) List(ctx context.Context, in model.RoleListInput) (out *model.RoleListOutput, err error) {
	var (
		m = dao.Role.Ctx(ctx)
	)
	out = &model.RoleListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	m = m.Where(dao.Role.Columns().IfShow, consts.RoleIfShowOk)
	if in.Search != "" {
		searchKeyword := "%" + in.Search + "%"
		m = m.WhereLike(dao.Role.Columns().Name, searchKeyword).WhereOrLike(dao.Role.Columns().Nickname, searchKeyword)
	}
	listModel := m.Page(in.Page, in.Size)
	if err = listModel.Scan(&out.List); err != nil {
		return
	}
	if len(out.List) == 0 {
		return
	}
	for i, item := range out.List {
		out.List[i].GenderText = consts.RoleGenderState(item.Gender).String()
	}
	out.Total, err = m.Count()
	if err != nil {
		return
	}
	return
}

// Detail 详情
func (*sRole) Detail(ctx context.Context, in model.RoleDetailInput) (out *model.RoleDetailOutput, err error) {
	var m = dao.Role.Ctx(ctx)
	m = m.Where(dao.Role.Columns().IfShow, consts.RoleIfShowOk)
	err = m.Where(dao.Role.Columns().Id, in.Id).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		err = errors.New(consts.RoleNotFoundErrMsg)
		return
	}
	if out.Id > 0 {
		out.GenderText = consts.RoleGenderState(out.Gender).String()
	}

	return
}
