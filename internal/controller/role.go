package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "threebody/api/v1"
	"threebody/internal/model"
	"threebody/internal/service"
)

type role struct{}

var Role = role{}

func (*role) Create(ctx context.Context, req *v1.RoleCreateReq) (res *v1.RoleRes, err error) {
	var input model.RoleCreateInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	err = service.Role().Create(ctx, input)
	return
}

func (*role) Update(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleRes, err error) {
	g.Dump(req.RoleId)
	var input model.RoleUpdateInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return nil, err
	}
	g.Dump(input)
	return
}
