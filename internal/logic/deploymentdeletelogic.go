package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentDeleteLogic {
	return DeploymentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentDeleteLogic) DeploymentDelete(req types.DeploymentDeleteReq) (*types.DeploymentDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentDeleteResp{}, nil
}
