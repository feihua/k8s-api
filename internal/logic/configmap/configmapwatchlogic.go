package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigMapWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapWatchLogic {
	return ConfigMapWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapWatchLogic) ConfigMapWatch(req types.ConfigMapWatchReq) (*types.ConfigMapWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigMapWatchResp{}, nil
}
