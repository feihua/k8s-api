package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigMapUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapUpdateStatusLogic {
	return ConfigMapUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapUpdateStatusLogic) ConfigMapUpdateStatus(req types.ConfigMapUpdateStatusReq) (*types.ConfigMapUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigMapUpdateStatusResp{}, nil
}
