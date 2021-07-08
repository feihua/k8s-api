package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceWatchLogic {
	return NamespaceWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceWatchLogic) NamespaceWatch(req types.NamespaceWatchReq) (*types.NamespaceWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespaceWatchResp{}, nil
}
