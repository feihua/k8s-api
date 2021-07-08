package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
