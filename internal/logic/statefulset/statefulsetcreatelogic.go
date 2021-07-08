package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
