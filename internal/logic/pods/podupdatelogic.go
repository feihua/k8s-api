package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

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
