package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
