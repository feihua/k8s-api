package logic

import (
	"context"
	"fmt"
	appsV1 "k8s.io/api/apps/v1beta1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/utils"
	"log"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeploymentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentAddLogic {
	return DeploymentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentAddLogic) DeploymentAdd(req types.DeploymentAddReq) (*types.DeploymentAddResp, error) {
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

	// 3. deployment 创建
	deploymentClient := clientset.AppsV1().Deployments("default")
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "test-nginx-dev1",
			Labels: map[string]string{
				"source": "cmdb",
				"app":    "nginx",
				"env":    "test",
			},
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: utils.Int32Ptr(3),
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"source": "cmdb",
					"app":    "nginx",
					"env":    "test",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"source": "cmdb",
						"app":    "nginx",
						"env":    "test",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: 80,
									Protocol:      coreV1.ProtocolTCP,
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Println("create deployment:")
	//deployment, err = deploymentClient.Create(context.TODO(), deployment, metaV1.CreateOptions{})
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println(deployment.Status)
	//}

	return &types.DeploymentAddResp{}, nil
}
