package jenkins

import (
	"context"
	"fmt"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBuildLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBuildLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBuildLogic {
	return GetBuildLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBuildLogic) GetBuild(req types.GetBuildReq) (*types.GetBuildResp, error) {
	build, err := l.svcCtx.Jenkins.GetBuild(l.ctx, req.JobName, req.Number)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询job build信息异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println(build)

	return &types.GetBuildResp{
		Code:    0,
		Message: "successful",
	}, nil

}
