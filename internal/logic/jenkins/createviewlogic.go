package jenkins

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateViewLogic {
	return CreateViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateViewLogic) CreateView(req types.CreateViewReq) (*types.CreateViewResp, error) {
	jobs, err := l.svcCtx.Jenkins.CreateView(l.ctx, req.Name, gojenkins.LIST_VIEW)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息失败,请求参数:%s,异常:%s", req.Name, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println(jobs)

	return &types.CreateViewResp{
		Code:    0,
		Message: "successful",
	}, nil

}
