package jenkins

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

	// Create jenkins job
	configString := `<?xml version='1.0' encoding='UTF-8'?>
						<project>
						  <actions/>
						  <description></description>
						  <keepDependencies>false</keepDependencies>
						  <properties/>
						  <scm class="hudson.scm.NullSCM"/>
						  <canRoam>true</canRoam>
						  <disabled>false</disabled>
						  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
						  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
						  <triggers class="vector"/>
						  <concurrentBuild>false</concurrentBuild>
						  <builders/>
						  <publishers/>
						  <buildWrappers/>
						</project>`
	job, err := l.svcCtx.Jenkins.CreateJob(l.ctx, configString, req.JobName)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("添加job失败,请求参数:%s,异常:%s", req.Config, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	view, err := l.svcCtx.Jenkins.GetView(l.ctx, req.ViewName)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("获取view失败,请求参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	b, err := view.AddJob(l.ctx, req.JobName)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("往view添加job失败,请求参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	if !b {
		logx.WithContext(l.ctx).Errorf("往view添加job失败,请求参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var result types.Result
	raw, _ := json.Marshal(job.Raw)
	json.Unmarshal(raw, &result)

	return &types.CreateJobResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data:    result,
	}, nil

}
