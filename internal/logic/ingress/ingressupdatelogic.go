package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
