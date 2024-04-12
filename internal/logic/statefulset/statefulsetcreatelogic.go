package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatefulSetCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetCreateLogic {
	return StatefulSetCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetCreateLogic) StatefulSetCreate(req types.StatefulSetAddReq) (*types.StatefulSetAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetAddResp{}, nil
}
