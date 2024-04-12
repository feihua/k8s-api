package jenkins

import (
	"context"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJobToViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddJobToViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddJobToViewLogic {
	return AddJobToViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddJobToViewLogic) AddJobToView(req types.AddJobToViewReq) (*types.AddJobToViewResp, error) {
	view, err := l.svcCtx.Jenkins.GetView(l.ctx, req.ViewName)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("获取view失败,请求参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	job, err := view.AddJob(l.ctx, req.JobName)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("往view添加job失败,请求参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	if !job {
		logx.WithContext(l.ctx).Errorf("往view添加job失败,请求参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.AddJobToViewResp{}, nil
}
