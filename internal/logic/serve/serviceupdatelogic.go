package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceUpdateLogic {
	return ServiceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceUpdateLogic) ServiceUpdate(req types.ServiceUpdateReq) (*types.ServiceUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.ServiceUpdateResp{}, nil
}
