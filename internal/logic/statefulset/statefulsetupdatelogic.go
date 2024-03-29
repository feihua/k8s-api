package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StatefulSetUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetUpdateLogic {
	return StatefulSetUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetUpdateLogic) StatefulSetUpdate(req types.StatefulSetUpdateReq) (*types.StatefulSetUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetUpdateResp{}, nil
}
