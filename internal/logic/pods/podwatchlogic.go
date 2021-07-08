package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PodWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodWatchLogic {
	return PodWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodWatchLogic) PodWatch(req types.PodWatchReq) (*types.PodWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodWatchResp{}, nil
}
