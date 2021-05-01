package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressListLogic {
	return IngressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressListLogic) IngressList(req types.IngressListReq) (*types.IngressListResp, error) {
	// todo: add your logic here and delete this line

	return &types.IngressListResp{}, nil
}
