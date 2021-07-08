package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StatefulSetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetListLogic {
	return StatefulSetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetListLogic) StatefulSetList(req types.StatefulSetListReq) (*types.StatefulSetListResp, error) {
	// todo: add your logic here and delete this line

	return &types.StatefulSetListResp{}, nil
}
