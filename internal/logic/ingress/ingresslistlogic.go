package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
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

	forConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	ingressClient := forConfig.ExtensionsV1beta1().Ingresses(req.Namespace)
	ingressList, err := ingressClient.List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.IngressListData
	for _, ingress := range ingressList.Items {
		list = append(list, &types.IngressListData{
			Name:              ingress.Name,
			Namespace:         ingress.Namespace,
			Host:              ingress.Spec.Rules[0].Host,
			ServiceName:       ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName,
			ServicePort:       ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServicePort.IntVal,
			CreationTimestamp: ingress.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})

	}

	return &types.IngressListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil
}
