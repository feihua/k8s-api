package jenkins

import (
	"context"
	"fmt"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteJobLogic {
	return DeleteJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteJobLogic) DeleteJob(req types.DeleteJenkinsJobReq) (*types.DeleteJenkinsJobResp, error) {
	jobs, err := l.svcCtx.Jenkins.DeleteJob(l.ctx, req.Name)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息失败,请求参数:%s,异常:%s", req.Name, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println(jobs)

	return &types.DeleteJenkinsJobResp{
		Code:    0,
		Message: "successful",
	}, nil

}
