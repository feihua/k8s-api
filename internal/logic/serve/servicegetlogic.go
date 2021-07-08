package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceGetLogic {
	return ServiceGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceGetLogic) ServiceGet(req types.ServiceGetReq) (*types.ServiceGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.ServiceGetResp{}, nil
}
