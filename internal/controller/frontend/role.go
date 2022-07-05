package frontend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
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
	res.Data, err = service.Role().List(ctx, input)
	g.Dump(res)
	os.Exit(-1)
	if err != nil {
		return nil, err
	}
	return
}
