package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentPatchLogic {
	return DeploymentPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentPatchLogic) DeploymentPatch(req types.DeploymentPatchReq) (*types.DeploymentPatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentPatchResp{}, nil
}
