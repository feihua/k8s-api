package logic

import (
	"context"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceCreateLogic {
	return ServiceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceCreateLogic) ServiceCreate(req types.ServiceAddReq) (*types.ServiceAddResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
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

	return &types.ServiceAddResp{}, nil
}
