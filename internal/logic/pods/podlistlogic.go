package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PodListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodListLogic {
	return PodListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodListLogic) PodList(req types.PodsListReq) (*types.PodsListResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodsListResp{}, nil
}
