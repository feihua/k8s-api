package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigMapUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapUpdateLogic {
	return ConfigMapUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapUpdateLogic) ConfigMapUpdate(req types.ConfigMapUpdateReq) (*types.ConfigMapUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigMapUpdateResp{}, nil
}
