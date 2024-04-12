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

type PodGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodGetLogic {
	return PodGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodGetLogic) PodGet(req types.PodGetReq) (*types.PodGetResp, error) {
	client := l.svcCtx.ClientSet.CoreV1().Pods(req.Namespace)
	pod, err := client.Get(context.TODO(), req.Pod, metaV1.GetOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询单个pod信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}
	data := types.PodsListItem{
		Name:               pod.Name,
		Status:             string(pod.Status.Phase),
		Labels:             pod.Labels,
		NodeSelector:       nil,
		Namespace:          pod.Namespace,
		HostIP:             pod.Status.HostIP,
		PodIP:              pod.Status.PodIP,
		StartTime:          pod.Status.StartTime.Format("2006-01-02 15:04:05"),
		RestartCount:       pod.Status.ContainerStatuses[0].RestartCount,
		Image:              pod.Status.ContainerStatuses[0].Image,
		CreationTimestamp:  pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		RestartPolicy:      string(pod.Spec.RestartPolicy),
		ServiceAccountName: pod.Spec.ServiceAccountName,
		NodeName:           pod.Spec.NodeName,
		HostNetwork:        pod.Spec.HostNetwork,
		//ImagePullSecrets:   pod.Spec.ImagePullSecrets[0].Name,
		Hostname: pod.Spec.Hostname,
	}

	dataStr, _ := json.Marshal(data)
	logx.WithContext(l.ctx).Infof("查询单个pod信息,请求参数：%s,响应：%s", req.Namespace, dataStr)

	return &types.PodGetResp{
		Code:    0,
		Message: "successful",
		Data:    data,
	}, nil

}
