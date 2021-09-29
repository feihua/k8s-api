package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigMapPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapPatchLogic {
	return ConfigMapPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapPatchLogic) ConfigMapPatch(req types.ConfigMapWatchReq) (*types.ConfigMapPatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.ConfigMapPatchResp{}, nil
}
