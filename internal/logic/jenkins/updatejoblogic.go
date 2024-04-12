package jenkins

import (
	"context"
	"fmt"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateJobLogic {
	return UpdateJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateJobLogic) UpdateJob(req types.UpdateJenkinsJobReq) (*types.UpdateJenkinsJobResp, error) {
	job := l.svcCtx.Jenkins.UpdateJob(l.ctx, req.Job, req.Config)

	fmt.Println(job)
	return &types.UpdateJenkinsJobResp{
		Code:    0,
		Message: "successful",
	}, nil

}
