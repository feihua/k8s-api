package jenkins

import (
	"context"
	"fmt"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateJobLogic {
	return CreateJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateJobLogic) CreateJob(req types.CreateJobReq) (*types.CreateJobResp, error) {
	jobs, err := l.svcCtx.Jenkins.CreateJob(l.ctx, req.Config)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息失败,请求参数:%s,异常:%s", req.Config, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println(jobs)

	return &types.CreateJobResp{
		Code:    0,
		Message: "successful",
	}, nil

}
