package logic

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/utils"
	"log"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentUpdateLogic {
	return DeploymentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentUpdateLogic) DeploymentUpdate(req types.DeploymentUpdateReq) (*types.DeploymentUpdateResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	deploymentClient := clientset.AppsV1().Deployments("default")

	//// 4. deployment 修改
	deployment, err := deploymentClient.Get(context.TODO(), "nginx-dev", metaV1.GetOptions{})

	if *deployment.Spec.Replicas > 3 {
		deployment.Spec.Replicas = utils.Int32Ptr(1)
	} else {
		deployment.Spec.Replicas = utils.Int32Ptr(*deployment.Spec.Replicas + 1)
	}
	// 1 => nginx:1.19.1
	// 2 => nginx:1.19.2
	// 3 => nginx:1.19.3
	// 3 => nginx:1.19.4
	deployment.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("nginx:1.19.%d", *deployment.Spec.Replicas)

	deployment, err = deploymentClient.Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	if err != nil {
		log.Println(err)
	}
	return &types.DeploymentUpdateResp{}, nil
}
