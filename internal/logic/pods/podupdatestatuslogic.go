package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
