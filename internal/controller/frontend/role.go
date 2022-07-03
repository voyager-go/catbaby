package frontend

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "threebody/api/v1"
	"threebody/internal/model"
)

type role struct{}

var Role = role{}

func (*role) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	var input model.RoleListInput
	err = gconv.Struct(req, &input)
	if err != nil {
		return
	}
	//res, err = service.Role().List(ctx, input)
	return
}
