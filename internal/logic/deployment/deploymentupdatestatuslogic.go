package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeploymentUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentUpdateStatusLogic {
	return DeploymentUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentUpdateStatusLogic) DeploymentUpdateStatus(req types.DeploymentUpdateStatusReq) (*types.DeploymentUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentUpdateStatusResp{}, nil
}
