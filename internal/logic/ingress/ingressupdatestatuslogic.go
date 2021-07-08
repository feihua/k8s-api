package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressUpdateStatusLogic {
	return IngressUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressUpdateStatusLogic) IngressUpdateStatus(req types.IngressUpdateStatusReq) (*types.IngressUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.IngressUpdateStatusResp{}, nil
}
