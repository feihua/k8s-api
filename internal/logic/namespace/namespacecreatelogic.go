package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceCreateLogic {
	return NamespaceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceCreateLogic) NamespaceCreate(req types.NamespaceAddReq) (*types.NamespaceAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespaceAddResp{}, nil
}
