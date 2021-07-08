package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentWatchLogic {
	return DeploymentWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentWatchLogic) DeploymentWatch(req types.DeploymentWatchReq) (*types.DeploymentWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentWatchResp{}, nil
}
