package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceDeleteLogic {
	return NamespaceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceDeleteLogic) NamespaceDelete(req types.NamespaceDeleteReq) (*types.NamespaceDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespaceDeleteResp{}, nil
}
