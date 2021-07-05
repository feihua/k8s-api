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
	"strings"
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
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	// 2. deployment 列表
	fmt.Println("deployments:")

	var list []*types.DeploymentListData

	if len(strings.TrimSpace(req.Namespace)) == 0 {
		namespaceClient := clientSet.CoreV1().Namespaces()
		namespaceResult, err := namespaceClient.List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			return nil, errorx.NewDefaultError(err.Error())
		}
		for _, namespace := range namespaceResult.Items {
			deploymentClient := clientSet.AppsV1().Deployments(namespace.Name)
			deploymentResult, err := deploymentClient.List(context.TODO(), metaV1.ListOptions{})

			if err != nil {
				return nil, errorx.NewDefaultError(err.Error())
			}

			for _, deployment := range deploymentResult.Items {
				fmt.Println(deployment.Name, deployment.Namespace, deployment.CreationTimestamp)
				list = append(list, &types.DeploymentListData{
					Name:               deployment.Name,
					Namespace:          deployment.Namespace,
					Labels:             deployment.Labels["run"],
					Strategy:           string(deployment.Spec.Strategy.Type),
					Replicas:           deployment.Status.Replicas,
					UpdatedReplicas:    deployment.Status.UpdatedReplicas,
					ReadyReplicas:      deployment.Status.ReadyReplicas,
					AvailableReplicas:  deployment.Status.AvailableReplicas,
					ObservedGeneration: deployment.Status.ObservedGeneration,
					CreationTimestamp:  deployment.CreationTimestamp.Format("2006-01-02 15:04:05"),
				})
			}
		}
	} else {
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
				Labels:             deployment.Labels["run"],
				Strategy:           string(deployment.Spec.Strategy.Type),
				Replicas:           deployment.Status.Replicas,
				UpdatedReplicas:    deployment.Status.UpdatedReplicas,
				ReadyReplicas:      deployment.Status.ReadyReplicas,
				AvailableReplicas:  deployment.Status.AvailableReplicas,
				ObservedGeneration: deployment.Status.ObservedGeneration,
				CreationTimestamp:  deployment.CreationTimestamp.Format("2006-01-02 15:04:05"),
			})
		}
	}

	return &types.DeploymentListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil
}
