package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceAddLogic {
	return NamespaceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceAddLogic) NamespaceAdd(req types.NamespaceAddReq) (*types.NamespaceAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespaceAddResp{}, nil
}
