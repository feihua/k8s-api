package logic

import (
	"context"
	"encoding/json"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	ingress, err := l.svcCtx.ClientSet.ExtensionsV1beta1().Ingresses(req.Namespace).Get(context.TODO(), req.Ingress, metaV1.GetOptions{})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("查询单个ingress信息失败,请求参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}
	rules, _ := json.Marshal(ingress.Spec.Rules)
	address, _ := json.Marshal(ingress.Status.LoadBalancer.Ingress)

	data := types.IngressListItem{
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
	}

	reqStr, _ := json.Marshal(req)
	dataStr, _ := json.Marshal(data)
	logx.WithContext(l.ctx).Infof("查询单个ingress信息,请求参数：%s,响应：%s", reqStr, dataStr)
	return &types.IngressGetResp{
		Code:    0,
		Message: "successful",
		Data:    data,
	}, nil
}
