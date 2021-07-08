package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
	"log"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceDeleteLogic {
	return NamespaceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceDeleteLogic) NamespaceDelete(req types.NamespaceDeleteReq) (*types.NamespaceDeleteResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		log.Fatal(err)
	}

	forConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	namespaceClient := forConfig.CoreV1().Namespaces()
	_ = namespaceClient.Delete(context.TODO(), req.Name, metaV1.DeleteOptions{})

	return &types.NamespaceDeleteResp{}, nil
}
