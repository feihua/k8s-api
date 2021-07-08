package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressGetLogic {
	return IngressGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressGetLogic) IngressGet(req types.IngressGetReq) (*types.IngressGetResp, error) {
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	forConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	ingress, err := forConfig.ExtensionsV1beta1().Ingresses(req.Namespace).Get(context.TODO(), req.Ingress, metaV1.GetOptions{})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.IngressGetResp{
		Code: 0,
		Msg:  "successful",
		Data: types.IngressGetData{
			Name:              ingress.Name,
			Namespace:         ingress.Namespace,
			Host:              ingress.Spec.Rules[0].Host,
			ServiceName:       ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName,
			ServicePort:       ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServicePort.IntVal,
			CreationTimestamp: ingress.CreationTimestamp.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
