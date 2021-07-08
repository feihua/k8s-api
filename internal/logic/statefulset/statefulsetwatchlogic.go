package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StatefulSetWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetWatchLogic {
	return StatefulSetWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetWatchLogic) StatefulSetWatch(req types.StatefulSetWatchReq) (*types.StatefulSetWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetWatchResp{}, nil
}
