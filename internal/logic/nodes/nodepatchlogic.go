package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodePatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodePatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodePatchLogic {
	return NodePatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodePatchLogic) NodePatch(req types.NodePatchReq) (*types.NodePatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodePatchResp{}, nil
}
