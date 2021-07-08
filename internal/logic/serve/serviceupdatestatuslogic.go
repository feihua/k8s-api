package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceUpdateStatusLogic {
	return ServiceUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceUpdateStatusLogic) ServiceUpdateStatus(req types.ServiceUpdateStatusReq) (*types.ServiceUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.ServiceUpdateStatusResp{}, nil
}
