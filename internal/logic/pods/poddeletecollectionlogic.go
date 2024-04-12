package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

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
