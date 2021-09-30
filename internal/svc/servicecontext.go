package svc

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (sc ServiceContext) NodeList() (*v1.NodeList, error) {
	return sc.ClientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
}
func (sc ServiceContext) NodeGet(name string) (*v1.Node, error) {
	return sc.ClientSet.CoreV1().Nodes().Get(context.TODO(), name, metaV1.GetOptions{})
}
