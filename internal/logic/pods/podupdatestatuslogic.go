package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodUpdateStatusLogic {
	return PodUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodUpdateStatusLogic) PodUpdateStatus(req types.PodUpdateStatusReq) (*types.PodUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodUpdateStatusResp{}, nil
}
