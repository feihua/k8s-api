package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeploymentGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentGetLogic {
	return DeploymentGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentGetLogic) DeploymentGet(req types.DeploymentGetReq) (*types.DeploymentGetResp, error) {
	client := l.svcCtx.ClientSet.AppsV1().Deployments(req.Namespace)
	result, err := client.Get(context.TODO(), req.Deployment, metaV1.GetOptions{})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("查询单个deployment信息失败,请求参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	data := types.DeploymentListItem{
		Name:               result.Name,
		Namespace:          result.Namespace,
		Labels:             result.Labels,
		Strategy:           string(result.Spec.Strategy.Type),
		Replicas:           result.Status.Replicas,
		UpdatedReplicas:    result.Status.UpdatedReplicas,
		ReadyReplicas:      result.Status.ReadyReplicas,
		AvailableReplicas:  result.Status.AvailableReplicas,
		ObservedGeneration: result.Status.ObservedGeneration,
		CreationTimestamp:  result.CreationTimestamp.Format("2006-01-02 15:04:05"),
		Images:             result.Spec.Template.Spec.Containers[0].Image,
		ImagePullPolicy:    string(result.Spec.Template.Spec.Containers[0].ImagePullPolicy),
		Message:            result.Status.Conditions[0].Message,
		Reason:             result.Status.Conditions[0].Reason,
		Status:             string(result.Status.Conditions[0].Status),
		LastUpdateTime:     result.Status.Conditions[0].LastUpdateTime.Format("2006-01-02 15:04:05"),
	}

	reqStr, _ := json.Marshal(req)
	dataStr, _ := json.Marshal(data)
	logx.WithContext(l.ctx).Infof("查询单个deployment信息,请求参数：%s,响应：%s", reqStr, dataStr)
	return &types.DeploymentGetResp{
		Code:    0,
		Message: "successful",
		Data:    data,
	}, nil

}
