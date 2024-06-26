package logic

import (
	"context"
	"github.com/feihua/k8s-api/internal/common/errorx"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
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

	client := l.svcCtx.ClientSet.ExtensionsV1beta1().Ingresses(req.Namespace)
	ingressList, err := client.List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询ingress列表信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.IngressListItem
	for _, ingress := range ingressList.Items {
		rules, _ := json.Marshal(ingress.Spec.Rules)
		address, _ := json.Marshal(ingress.Status.LoadBalancer.Ingress)
		list = append(list, &types.IngressListItem{
			Name:              ingress.Name,
			Namespace:         ingress.Namespace,
			Host:              ingress.Spec.Rules[0].Host,
			ServiceName:       ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServiceName,
			ServicePort:       ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths[0].Backend.ServicePort.IntVal,
			CreationTimestamp: ingress.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Labels:            ingress.Labels,
			Status:            ingress.Status.String(),
			Rules:             string(rules),
			Address:           string(address),
			Annotations:       ingress.Annotations,
			ResourceVersion:   ingress.ResourceVersion,
		})

	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询ingress列表信息,请求参数：%s,响应：%s", req.Namespace, listStr)
	return &types.IngressListResp{
		Code:    0,
		Message: "successful",
		Data: types.IngressListData{
			Items: list,
			Total: int64(len(list)),
		},
	}, nil
}
