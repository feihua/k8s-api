package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodUpdateEphemeralContainersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodUpdateEphemeralContainersLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodUpdateEphemeralContainersLogic {
	return PodUpdateEphemeralContainersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodUpdateEphemeralContainersLogic) PodUpdateEphemeralContainers(req types.PodUpdateEphemeralContainersReq) (*types.PodUpdateEphemeralContainersResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodUpdateEphemeralContainersResp{}, nil
}
