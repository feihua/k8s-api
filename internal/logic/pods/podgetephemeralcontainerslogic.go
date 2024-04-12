package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodGetEphemeralContainersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodGetEphemeralContainersLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodGetEphemeralContainersLogic {
	return PodGetEphemeralContainersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodGetEphemeralContainersLogic) PodGetEphemeralContainers(req types.PodGetEphemeralContainersReq) (*types.PodGetEphemeralContainersResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodGetEphemeralContainersResp{}, nil
}
