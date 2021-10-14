package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleListLogic {
	return RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req types.ListRoleReq) (*types.ListRoleResp, error) {
	// todo: add your logic here and delete this line

	items := make([]*types.ListRoleItem, 2)

	items[0] = &types.ListRoleItem{
		Id:         1,
		OrderNo:    "1",
		RoleName:   "超级管理员",
		RoleValue:  "admin",
		CreateTime: "2021-10-11",
		Status:     "1",
		Remark:     "就是牛逼，牛哄哄的",
	}
	items[1] = &types.ListRoleItem{
		Id:         2,
		OrderNo:    "2",
		RoleName:   "普通管理员",
		RoleValue:  "manager",
		CreateTime: "2021-10-11",
		Status:     "1",
		Remark:     "实力一半，哈哈",
	}
	return &types.ListRoleResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.ListRoleData{
			Items: items,
			Total: 2,
		},
	}, nil
}
