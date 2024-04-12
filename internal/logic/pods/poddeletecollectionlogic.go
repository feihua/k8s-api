package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodDeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodDeleteCollectionLogic {
	return PodDeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodDeleteCollectionLogic) PodDeleteCollection(req types.PodDeleteCollectionReq) (*types.PodDeleteCollectionResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodDeleteCollectionResp{}, nil
}
