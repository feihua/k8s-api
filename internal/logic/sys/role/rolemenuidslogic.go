package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleMenuIdsLogic {
	return RoleMenuIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleMenuIdsLogic) RoleMenuIds(req types.RoleMenuIdsReq) (*types.RoleMenuIdsResp, error) {
	// todo: add your logic here and delete this line

	return &types.RoleMenuIdsResp{}, nil
}
