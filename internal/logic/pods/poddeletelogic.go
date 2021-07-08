package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PodDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodDeleteLogic {
	return PodDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodDeleteLogic) PodDelete(req types.PodDeleteReq) (*types.PodDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodDeleteResp{}, nil
}
