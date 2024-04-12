package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatefulSetDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetDeleteLogic {
	return StatefulSetDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetDeleteLogic) StatefulSetDelete(req types.StatefulSetDeleteReq) (*types.StatefulSetDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetDeleteResp{}, nil
}
