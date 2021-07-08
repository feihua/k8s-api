package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) PatchLogic {
	return PatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchLogic) Patch(req types.NamespaceListReq) (*types.NamespacePatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.NamespacePatchResp{}, nil
}
