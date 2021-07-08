package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeDeleteLogic {
	return NodeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeDeleteLogic) NodeDelete(req types.NodeDeleteReq) (*types.NodeDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodeDeleteResp{}, nil
}
