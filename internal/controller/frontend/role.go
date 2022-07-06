package frontend

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "threebody/api/v1"
	"threebody/internal/model"
	"threebody/internal/service"
)

type role struct{}

var Role = role{}

// List 故事角色列表
func (*role) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	var input model.RoleListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return
	}
	list, err := service.Role().List(ctx, input)
	res = &v1.RoleListRes{
		RoleListOutput: list,
	}
	if err != nil {
		return
	}
	return
}

// Detail 故事角色详情
func (*role) Detail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	var input = model.RoleDetailInput{Id: gconv.Uint(req.Id)}
	detail, err := service.Role().Detail(ctx, input)
	res = &v1.RoleDetailRes{
		RoleDetailOutput: detail,
	}
	return
}
