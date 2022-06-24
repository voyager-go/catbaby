package service

import (
	"context"
	"threebody/internal/model"
)

type IRole interface {
	Create(ctx context.Context, in model.RoleCreateInput) error
}

var iRole IRole

func Role() IRole {
	if iRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return iRole
}

func RegisterRole(i IRole) {
	iRole = i
}
