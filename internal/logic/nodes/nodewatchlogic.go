package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeWatchLogic {
	return NodeWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeWatchLogic) NodeWatch(req types.NodeWatchReq) (*types.NodeWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodeWatchResp{}, nil
}
