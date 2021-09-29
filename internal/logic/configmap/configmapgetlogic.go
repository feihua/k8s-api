package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigMapGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapGetLogic {
	return ConfigMapGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapGetLogic) ConfigMapGet(req types.ConfigMapGetReq) (*types.ConfigMapGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigMapGetResp{}, nil
}
