package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeploymentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentCreateLogic {
	return DeploymentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentCreateLogic) DeploymentCreate(req types.DeploymentAddReq) (*types.DeploymentAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentAddResp{}, nil
}
