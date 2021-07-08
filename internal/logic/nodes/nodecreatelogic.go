package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeCreateLogic {
	return NodeCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeCreateLogic) NodeCreate(req types.NodeAddReq) (*types.NodeAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodeAddResp{}, nil
}
