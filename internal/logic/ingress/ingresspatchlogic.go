package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressPatchLogic {
	return IngressPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressPatchLogic) IngressPatch(req types.IngressPatchReq) (*types.IngressPatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.IngressPatchResp{}, nil
}
