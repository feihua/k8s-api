package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeListLogic {
	return NodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeListLogic) NodeList(req types.NodesListReq) (*types.NodesListResp, error) {
	// todo: add your logic here and delete this line

	return &types.NodesListResp{}, nil
}
