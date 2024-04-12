package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodUpdateLogic {
	return PodUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodUpdateLogic) PodUpdate(req types.PodUpdateReq) (*types.PodUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodUpdateResp{}, nil
}
