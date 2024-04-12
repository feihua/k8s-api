package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentListLogic {
	return DeploymentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentListLogic) DeploymentList(req types.DeploymentListReq) (*types.DeploymentListResp, error) {
	client := l.svcCtx.ClientSet.AppsV1().Deployments(req.Namespace)
	result, err := client.List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询deployment列表信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.DeploymentListItem

	for _, item := range result.Items {
		list = append(list, &types.DeploymentListItem{
			Name:               item.Name,
			Namespace:          item.Namespace,
			Labels:             item.Labels,
			Strategy:           string(item.Spec.Strategy.Type),
			Replicas:           item.Status.Replicas,
			UpdatedReplicas:    item.Status.UpdatedReplicas,
			ReadyReplicas:      item.Status.ReadyReplicas,
			AvailableReplicas:  item.Status.AvailableReplicas,
			ObservedGeneration: item.Status.ObservedGeneration,
			CreationTimestamp:  item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Images:             item.Spec.Template.Spec.Containers[0].Image,
			ImagePullPolicy:    string(item.Spec.Template.Spec.Containers[0].ImagePullPolicy),
			Message:            item.Status.Conditions[0].Message,
			Reason:             item.Status.Conditions[0].Reason,
			Status:             string(item.Status.Conditions[0].Status),
			LastUpdateTime:     item.Status.Conditions[0].LastUpdateTime.Format("2006-01-02 15:04:05"),
		})

	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询deployment列表信息,请求参数：%s,响应：%s", req.Namespace, listStr)
	return &types.DeploymentListResp{
		Code:    0,
		Message: "successful",
		Data: types.DeploymentListData{
			Items: list,
			Total: int64(len(list)),
		},
	}, nil
}
