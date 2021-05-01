package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentListLogic {
	return DeploymentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentListLogic) DeploymentList(req types.DeploymentListReq) (*types.DeploymentListResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentListResp{}, nil
}
