package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentGetLogic {
	return DeploymentGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentGetLogic) DeploymentGet(req types.DeploymentGetReq) (*types.DeploymentGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentGetResp{}, nil
}
