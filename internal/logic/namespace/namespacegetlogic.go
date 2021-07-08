package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceGetLogic {
	return NamespaceGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceGetLogic) NamespaceGet(req types.NamespaceGetReq) (*types.NamespaceGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespaceGetResp{}, nil
}
