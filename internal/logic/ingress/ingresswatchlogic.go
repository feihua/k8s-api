package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressWatchLogic {
	return IngressWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressWatchLogic) IngressWatch(req types.IngressWatchReq) (*types.IngressWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.IngressWatchResp{}, nil
}
