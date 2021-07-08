package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeUpdateStatusLogic {
	return NodeUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeUpdateStatusLogic) NodeUpdateStatus(req types.NodeUpdateStatusReq) (*types.NodeUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodeUpdateStatusResp{}, nil
}
