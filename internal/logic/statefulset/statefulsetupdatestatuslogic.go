package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatefulSetUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetUpdateStatusLogic {
	return StatefulSetUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetUpdateStatusLogic) StatefulSetUpdateStatus(req types.StatefulSetUpdateStatusReq) (*types.StatefulSetUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetUpdateStatusResp{}, nil
}
