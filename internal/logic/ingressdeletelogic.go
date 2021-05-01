package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressDeleteLogic {
	return IngressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressDeleteLogic) IngressDelete(req types.IngressDeleteReq) (*types.IngressDeleteReq, error) {
	// todo: add your logic here and delete this line

	return &types.IngressDeleteReq{}, nil
}
