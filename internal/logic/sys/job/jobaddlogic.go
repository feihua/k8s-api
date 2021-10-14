package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type JobAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) JobAddLogic {
	return JobAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobAddLogic) JobAdd(req types.AddJobReq) (*types.AddJobResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddJobResp{}, nil
}
