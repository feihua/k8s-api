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

type ServiceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceDeleteLogic {
	return ServiceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceDeleteLogic) ServiceDelete(req types.ServiceDeleteReq) (*types.ServiceDeleteResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		log.Fatal(err)
	}

	newForConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	serviceClient := newForConfig.CoreV1().Services("default")
	_ = serviceClient.Delete(context.TODO(), "nginx", metaV1.DeleteOptions{})

	return &types.ServiceDeleteResp{}, nil
}
