package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigMapDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapDeleteLogic {
	return ConfigMapDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapDeleteLogic) ConfigMapDelete(req types.ConfigMapDeleteReq) (*types.ConfigMapDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigMapDeleteResp{}, nil
}
