package logic

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

type ServiceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceListLogic {
	return ServiceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceListLogic) ServiceList(req types.ServiceListReq) (*types.ServiceListResp, error) {

	client := l.svcCtx.ClientSet.CoreV1().Services(req.Namespace)
	result, err := client.List(context.TODO(), metaV1.ListOptions{})

	var list []*types.ServiceListData
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询service列表信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}
	for _, service := range result.Items {
		list = append(list, &types.ServiceListData{
			Name:              service.Name,
			Namespace:         service.Namespace,
			Labels:            service.Labels,
			Annotations:       service.Annotations,
			Selector:          service.Spec.Selector,
			Type:              string(service.Spec.Type),
			ClusterIP:         service.Spec.ClusterIP,
			Protocol:          string(service.Spec.Ports[0].Protocol),
			Ports:             service.Spec.Ports[0].Port,
			TargetPort:        service.Spec.Ports[0].TargetPort.IntVal,
			NodePort:          service.Spec.Ports[0].NodePort,
			CreationTimestamp: service.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询service列表信息,请求参数：%s,响应：%s", req.Namespace, listStr)
	return &types.ServiceListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil
}
