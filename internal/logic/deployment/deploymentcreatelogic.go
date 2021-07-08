package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
