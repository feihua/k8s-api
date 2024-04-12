package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IngressCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressCreateLogic {
	return IngressCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressCreateLogic) IngressCreate(req types.IngressAddReq) (*types.IngressAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.IngressAddResp{}, nil
}
