package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceUpdateStatusLogic {
	return NamespaceUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceUpdateStatusLogic) NamespaceUpdateStatus(req types.NamespaceUpdateStatusReq) (*types.NamespaceUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespaceUpdateStatusResp{}, nil
}
