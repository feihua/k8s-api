package svc

import (
	"github.com/tal-tech/go-zero/core/logx"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	ClientSet *kubernetes.Clientset
}

func NewServiceContext(c config.Config) *ServiceContext {

	kubeConfig := "etc/config"
	buildConfigFromFlags, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		logx.Errorf("获取kubeConfig：%s异常: %s", kubeConfig, err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(buildConfigFromFlags)
	if err != nil {
		logx.Errorf("连接kubernetes：%s异常: %s", kubeConfig, err.Error())
	}

	return &ServiceContext{
		Config:    c,
		ClientSet: clientSet,
	}
}
