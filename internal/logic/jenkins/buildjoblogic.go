package jenkins

import (
	"context"
	"fmt"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BuildJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBuildJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) BuildJobLogic {
	return BuildJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BuildJobLogic) BuildJob(req types.BuildJobReq) (*types.BuildJobResp, error) {
	jobs, err := l.svcCtx.Jenkins.BuildJob(l.ctx, req.Name, req.Params)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息失败,请求参数:%s,异常:%s", req.Name, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println(jobs)
	return &types.BuildJobResp{
		Code:    0,
		Message: "successful",
	}, nil

}
