package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type WatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) WatchLogic {
	return WatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WatchLogic) Watch(req types.NamespaceListReq) (*types.NamespaceWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespaceWatchResp{}, nil
}
