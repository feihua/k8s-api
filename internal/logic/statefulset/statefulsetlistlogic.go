package logic

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"log"
	"time"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StatefulSetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetListLogic {
	return StatefulSetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetListLogic) StatefulSetList(req types.StatefulSetListReq) (*types.StatefulSetListResp, error) {
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	ingressClient := clientset.AppsV1().StatefulSets("")
	namespaceResult, err := ingressClient.List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	now := time.Now()

	var namespaces []string
	fmt.Println("namespaces:")
	for _, namespace := range namespaceResult.Items {
		namespaces = append(namespaces, namespace.Name)
		fmt.Println(namespace.Name, now.Sub(namespace.CreationTimestamp.Time))
	}

	// 2. deployment 列表
	fmt.Println("deployments:")
	for _, namespace := range namespaces {
		deploymentClient := clientset.AppsV1().Deployments(namespace)

		depoymentResult, err := deploymentClient.List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			log.Println(err)
		} else {
			for _, deployment := range depoymentResult.Items {
				fmt.Println(deployment.Name, deployment.Namespace, deployment.CreationTimestamp)
			}

		}
	}
	// 3. deployment 创建
	//deploymentClient := clientset.AppsV1().Deployments("default")
	//deployment := &appsV1.Deployment{
	//	ObjectMeta: metaV1.ObjectMeta{
	//		Name: "test-nginx-dev1",
	//		Labels: map[string]string{
	//			"source": "cmdb",
	//			"app":    "nginx",
	//			"env":    "test",
	//		},
	//	},
	//	Spec: appsV1.DeploymentSpec{
	//		Replicas: utils.Int32Ptr(3),
	//		Selector: &metaV1.LabelSelector{
	//			MatchLabels: map[string]string{
	//				"source": "cmdb",
	//				"app":    "nginx",
	//				"env":    "test",
	//			},
	//		},
	//		Template: coreV1.PodTemplateSpec{
	//			ObjectMeta: metaV1.ObjectMeta{
	//				Labels: map[string]string{
	//					"source": "cmdb",
	//					"app":    "nginx",
	//					"env":    "test",
	//				},
	//			},
	//			Spec: coreV1.PodSpec{
	//				Containers: []coreV1.Container{
	//					{
	//						Name:  "nginx",
	//						Image: "nginx:latest",
	//						Ports: []coreV1.ContainerPort{
	//							{
	//								Name:          "http",
	//								ContainerPort: 80,
	//								Protocol:      coreV1.ProtocolTCP,
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//fmt.Println("create deployment:")
	//deployment, err = deploymentClient.Create(context.TODO(), deployment, metaV1.CreateOptions{})
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println(deployment.Status)
	//}
	//// 4. deployment 修改
	//deployment, err = deploymentClient.Get(context.TODO(), "nginx-dev", metaV1.GetOptions{})
	//
	//if *deployment.Spec.Replicas > 3 {
	//	deployment.Spec.Replicas = utils.Int32Ptr(1)
	//} else {
	//	deployment.Spec.Replicas = utils.Int32Ptr(*deployment.Spec.Replicas + 1)
	//}
	//// 1 => nginx:1.19.1
	//// 2 => nginx:1.19.2
	//// 3 => nginx:1.19.3
	//// 3 => nginx:1.19.4
	//deployment.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("nginx:1.19.%d", *deployment.Spec.Replicas)
	//
	//deployment, err = deploymentClient.Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	if err != nil {
		log.Println(err)
	}
	// 5. service 列表
	fmt.Println("services:")
	for _, namespace := range namespaces {
		serviceClient := clientset.CoreV1().Services(namespace)
		serviceResult, err := serviceClient.List(context.TODO(), metaV1.ListOptions{})
		if err != nil {
			log.Println(err)
		} else {
			for _, service := range serviceResult.Items {
				fmt.Println(service.Name, service.Namespace, service.Labels, service.Spec.Selector, service.Spec.Type, service.Spec.ClusterIP, service.Spec.Ports, service.CreationTimestamp)
			}
		}
	}
	// 6. service 创建
	serviceClient := clientset.CoreV1().Services("default")
	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "test-nginx-dev",
			Labels: map[string]string{
				"source": "cmdb",
				"app":    "nginx",
				"env":    "test",
			},
		},
		Spec: coreV1.ServiceSpec{
			Selector: map[string]string{
				"source": "cmdb",
				"app":    "nginx",
				"env":    "test",
			},
			Type: coreV1.ServiceTypeNodePort,
			Ports: []coreV1.ServicePort{
				{
					Name:     "http",
					Port:     80,
					Protocol: coreV1.ProtocolTCP,
				},
			},
		},
	}
	service, err = serviceClient.Create(context.TODO(), service, metaV1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
	}
	// 7. service 修改

	service, err = serviceClient.Get(context.TODO(), "nginx-dev", metaV1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	if service.Spec.Type == coreV1.ServiceTypeNodePort {
		service.Spec.Type = coreV1.ServiceTypeClusterIP
	} else {
		service.Spec.Type = coreV1.ServiceTypeNodePort
	}
	serviceClient.Update(context.TODO(), service, metaV1.UpdateOptions{})
	// 8. deployment 删除
	//deploymentClient.Delete(context.TODO(), "nginx", metaV1.DeleteOptions{})
	// 9. service 删除
	//serviceClient.Delete(context.TODO(), "nginx", metaV1.DeleteOptions{})

	return &types.StatefulSetListResp{}, nil
}
