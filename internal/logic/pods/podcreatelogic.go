package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodCreateLogic {
	return PodCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodCreateLogic) PodCreate(req types.PodAddReq) (*types.PodAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodAddResp{}, nil
}
