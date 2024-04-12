package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatefulSetGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetGetLogic {
	return StatefulSetGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetGetLogic) StatefulSetGet(req types.StatefulSetGetReq) (*types.StatefulSetGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetGetResp{}, nil
}
