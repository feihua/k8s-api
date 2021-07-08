package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StatefulSetPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetPatchLogic {
	return StatefulSetPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetPatchLogic) StatefulSetPatch(req types.StatefulSetPatchReq) (*types.StatefulSetPatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetPatchResp{}, nil
}
