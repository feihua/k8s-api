package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentDeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentDeleteCollectionLogic {
	return DeploymentDeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentDeleteCollectionLogic) DeploymentDeleteCollection(req types.DeploymentDeleteCollectionReq) (*types.DeploymentDeleteCollectionResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeploymentDeleteCollectionResp{}, nil
}
