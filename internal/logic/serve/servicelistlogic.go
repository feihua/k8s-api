package logic

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
	"log"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceListLogic {
	return ServiceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceListLogic) ServiceList(req types.ServiceListReq) (*types.ServiceListResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		log.Fatal(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// 5. service 列表
	fmt.Println("services:")
	//for _, namespace := range namespaces {
	//	serviceClient := clientSet.CoreV1().Services(namespace)
	//	serviceResult, err := serviceClient.List(context.TODO(), metaV1.ListOptions{})
	//	if err != nil {
	//		log.Println(err)
	//	} else {
	//		for _, service := range serviceResult.Items {
	//			fmt.Println(service.Name, service.Namespace, service.Labels, service.Spec.Selector, service.Spec.Type, service.Spec.ClusterIP, service.Spec.Ports, service.CreationTimestamp)
	//		}
	//	}
	//}

	serviceClient := clientSet.CoreV1().Services(req.Namespace)
	serviceResult, err := serviceClient.List(context.TODO(), metaV1.ListOptions{})

	var list []*types.ServiceListData
	if err != nil {
		log.Println(err)
	} else {
		for _, service := range serviceResult.Items {
			fmt.Println(service.Name, service.Namespace, service.Labels, service.Spec.Selector, service.Spec.Type, service.Spec.ClusterIP, service.Spec.Ports, service.CreationTimestamp)
			list = append(list, &types.ServiceListData{
				Name:      service.Name,
				Namespace: service.Namespace,
				//Labels:            service.Labels,
				//Selector:          service.Spec.Selector,
				//Type:              service.Spec.Type,
				//ClusterIP:         service.Spec.ClusterIP,
				//Ports:             service.Spec.Ports,
				CreationTimestamp: service.CreationTimestamp.Format("2006-01-02 15:04:05"),
			})
		}
	}

	return &types.ServiceListResp{
		ListData: list,
	}, nil
}