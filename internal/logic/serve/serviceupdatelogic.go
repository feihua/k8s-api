package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
