package logic

import (
	"context"
	"k8s_test/internal/common/errorx"
	"time"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"log"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type NamespaceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceListLogic {
	return NamespaceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceListLogic) NamespaceList(req types.NamespaceListReq) (*types.NamespaceListResp, error) {

	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		log.Fatal(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	// 1. namespace 列表
	namespaceClient := clientSet.CoreV1().Namespaces()
	namespaceResult, err := namespaceClient.List(context.TODO(), metaV1.ListOptions{})
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

	return &types.NamespaceListResp{
		Namespaces: namespaces,
	}, nil
}
