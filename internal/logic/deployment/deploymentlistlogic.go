package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
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
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.DeploymentListData
	deploymentClient := clientSet.AppsV1().Deployments(req.Namespace)
	deploymentResult, err := deploymentClient.List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	for _, deployment := range deploymentResult.Items {
		fmt.Println(deployment.Name, deployment.Namespace, deployment.CreationTimestamp)
		list = append(list, &types.DeploymentListData{
			Name:               deployment.Name,
			Namespace:          deployment.Namespace,
			Labels:             deployment.Labels,
			Strategy:           string(deployment.Spec.Strategy.Type),
			Replicas:           deployment.Status.Replicas,
			UpdatedReplicas:    deployment.Status.UpdatedReplicas,
			ReadyReplicas:      deployment.Status.ReadyReplicas,
			AvailableReplicas:  deployment.Status.AvailableReplicas,
			ObservedGeneration: deployment.Status.ObservedGeneration,
			CreationTimestamp:  deployment.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Images:             deployment.Spec.Template.Spec.Containers[0].Image,
			ImagePullPolicy:    string(deployment.Spec.Template.Spec.Containers[0].ImagePullPolicy),
			Message:            deployment.Status.Conditions[0].Message,
			Reason:             deployment.Status.Conditions[0].Reason,
			Status:             string(deployment.Status.Conditions[0].Status),
			LastUpdateTime:     deployment.Status.Conditions[0].LastUpdateTime.Format("2006-01-02 15:04:05"),
		})

		logx.Info("de%v", deployment)
	}

	return &types.DeploymentListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil
}
