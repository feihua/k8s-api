package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespacePatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespacePatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespacePatchLogic {
	return NamespacePatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespacePatchLogic) NamespacePatch(req types.NamespaceWatchReq) (*types.NamespacePatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespacePatchResp{}, nil
}
