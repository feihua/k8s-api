package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeUpdateLogic {
	return NodeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeUpdateLogic) NodeUpdate(req types.NodeUpdateReq) (*types.NodeUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodeUpdateResp{}, nil
}
