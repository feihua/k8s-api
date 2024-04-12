package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigMapCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapCreateLogic {
	return ConfigMapCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapCreateLogic) ConfigMapCreate(req types.ConfigMapAddReq) (*types.ConfigMapAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigMapAddResp{}, nil
}
