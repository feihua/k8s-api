package logic

import (
	"context"
	"encoding/json"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ServiceGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceGetLogic {
	return ServiceGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceGetLogic) ServiceGet(req types.ServiceGetReq) (*types.ServiceGetResp, error) {
	client := l.svcCtx.ClientSet.CoreV1().Services(req.Namespace)
	result, err := client.Get(context.TODO(), req.Service, metaV1.GetOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询单个service信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}
	data := types.ServiceListData{
		Name:              result.Name,
		Namespace:         result.Namespace,
		Labels:            result.Labels,
		Annotations:       result.Annotations,
		Selector:          result.Spec.Selector,
		Type:              string(result.Spec.Type),
		ClusterIP:         result.Spec.ClusterIP,
		Protocol:          string(result.Spec.Ports[0].Protocol),
		Ports:             result.Spec.Ports[0].Port,
		TargetPort:        result.Spec.Ports[0].TargetPort.IntVal,
		NodePort:          result.Spec.Ports[0].NodePort,
		CreationTimestamp: result.CreationTimestamp.Format("2006-01-02 15:04:05"),
	}

	dataStr, _ := json.Marshal(data)
	logx.WithContext(l.ctx).Infof("查询单个service信息,请求参数：%s,响应：%s", req.Namespace, dataStr)

	return &types.ServiceGetResp{
		Code: 0,
		Msg:  "successful",
		Data: data,
	}, nil

}
