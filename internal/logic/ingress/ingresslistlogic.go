package logic

import (
	"context"
	"fmt"
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

type IngressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressListLogic {
	return IngressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressListLogic) IngressList(req types.IngressListReq) (*types.IngressListResp, error) {

	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	ingressClient := clientset.ExtensionsV1beta1().Ingresses("")
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

	return &types.IngressListResp{}, nil
}
