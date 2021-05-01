package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentAddLogic {
	return DeploymentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentAddLogic) DeploymentAdd(req types.DeploymentAddReq) (*types.DeploymentAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentAddResp{}, nil
}
