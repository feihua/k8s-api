package logic

import (
	"context"
	"k8s_test/internal/common/errorx"
	"time"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
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
		return nil, errorx.NewDefaultError(err.Error())
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

	var listData []*types.NamespaceListData
	fmt.Println("namespace:")
	for _, namespace := range namespaceResult.Items {
		listData = append(listData, &types.NamespaceListData{
			Name:              namespace.Name,
			Status:            string(namespace.Status.Phase),
			CreationTimestamp: namespace.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})

		fmt.Println(namespace.Name, now.Sub(namespace.CreationTimestamp.Time))
	}

	return &types.NamespaceListResp{
		Code: 0,
		Msg:  "successful",
		Data: listData,
	}, nil
}
