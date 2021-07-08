package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceCreateLogic {
	return ServiceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceCreateLogic) ServiceCreate(req types.ServiceAddReq) (*types.ServiceAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.ServiceAddResp{}, nil
}
