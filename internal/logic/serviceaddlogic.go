package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceAddLogic {
	return ServiceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceAddLogic) ServiceAdd(req types.ServiceAddReq) (*types.ServiceAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.ServiceAddResp{}, nil
}
