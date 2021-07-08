package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeDeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeDeleteCollectionLogic {
	return NodeDeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeDeleteCollectionLogic) NodeDeleteCollection(req types.NodeDeleteCollectionReq) (*types.NodeDeleteCollectionResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodeDeleteCollectionResp{}, nil
}
