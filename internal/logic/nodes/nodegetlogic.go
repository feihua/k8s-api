package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeGetLogic {
	return NodeGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeGetLogic) NodeGet(req types.NodeGetReq) (*types.NodeGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodeGetResp{}, nil
}
