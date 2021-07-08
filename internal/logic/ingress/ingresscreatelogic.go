package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
