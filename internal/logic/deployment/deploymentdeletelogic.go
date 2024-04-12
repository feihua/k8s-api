package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

type DeploymentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentDeleteLogic {
	return DeploymentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentDeleteLogic) DeploymentDelete(req types.DeploymentDeleteReq) (*types.DeploymentDeleteResp, error) {

	deploymentClient := l.svcCtx.ClientSet.AppsV1().Deployments(req.Namespace)

	err := deploymentClient.Delete(context.TODO(), req.Deployment, metaV1.DeleteOptions{})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("查询单个deployment信息失败,请求参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.DeploymentDeleteResp{
		Code:    0,
		Message: "删除deployment: " + req.Deployment + "成功",
	}, nil
}
