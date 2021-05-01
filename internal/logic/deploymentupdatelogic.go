package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentUpdateLogic {
	return DeploymentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentUpdateLogic) DeploymentUpdate(req types.DeploymentUpdateReq) (*types.DeploymentUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentUpdateResp{}, nil
}
