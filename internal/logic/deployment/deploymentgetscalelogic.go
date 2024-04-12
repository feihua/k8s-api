package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeploymentGetScaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentGetScaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentGetScaleLogic {
	return DeploymentGetScaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentGetScaleLogic) DeploymentGetScale(req types.DeploymentGetScaleReq) (*types.DeploymentGetScaleResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentGetScaleResp{}, nil
}
