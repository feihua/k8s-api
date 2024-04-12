package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressUpdateLogic {
	return IngressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressUpdateLogic) IngressUpdate(req types.IngressUpdateReq) (*types.IngressUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.IngressUpdateResp{}, nil
}
