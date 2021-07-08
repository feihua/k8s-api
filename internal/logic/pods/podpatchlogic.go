package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PodPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodPatchLogic {
	return PodPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodPatchLogic) PodPatch(req types.PodPatchReq) (*types.PodPatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.PodPatchResp{}, nil
}
