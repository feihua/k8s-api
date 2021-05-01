package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressAddLogic {
	return IngressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressAddLogic) IngressAdd(req types.IngressAddReq) (*types.IngressAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.IngressAddResp{}, nil
}
