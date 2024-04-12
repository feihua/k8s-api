package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeploymentUpdateScaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentUpdateScaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentUpdateScaleLogic {
	return DeploymentUpdateScaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentUpdateScaleLogic) DeploymentUpdateScale(req types.DeploymentUpdateScaleReq) (*types.DeploymentUpdateScaleResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentUpdateScaleResp{}, nil
}
