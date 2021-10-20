package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PodLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodLogLogic {
	return PodLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodLogLogic) PodLog(req types.PodLogReq) (*types.PodLogResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodLogResp{}, nil
}
