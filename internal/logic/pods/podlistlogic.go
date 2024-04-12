package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodListLogic {
	return PodListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodListLogic) PodList(req types.PodsListReq) (*types.PodsListResp, error) {
	pods, err := l.svcCtx.ClientSet.CoreV1().Pods(req.Namespace).List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询pod列表信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.PodsListItem
	for _, pod := range pods.Items {
		list = append(list, &types.PodsListItem{
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
		})

	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询pod列表信息,请求参数：%s,响应：%s", req.Namespace, listStr)
	return &types.PodsListResp{
		Code:    0,
		Message: "successful",
		Data: types.PodsListData{
			Items: list,
			Total: int64(len(list)),
		},
	}, nil

}
