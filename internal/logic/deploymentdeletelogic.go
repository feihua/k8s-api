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

type DeploymentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeploymentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeploymentDeleteLogic {
	return DeploymentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeploymentDeleteLogic) DeploymentDelete(req types.DeploymentDeleteReq) (*types.DeploymentDeleteResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err2 := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err2 != nil {
		log.Fatal(err2)
	}
	newForConfig, err2 := kubernetes.NewForConfig(config)
	if err2 != nil {
		log.Fatal(err2)
	}

	deploymentClient := newForConfig.AppsV1().Deployments("default")

	deploymentClient.Delete(context.TODO(), "nginx", metaV1.DeleteOptions{})

	return &types.DeploymentDeleteResp{}, nil
}
