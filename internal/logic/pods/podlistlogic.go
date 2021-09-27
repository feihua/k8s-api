package logic

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.PodsListData
	podClient := clientSet.CoreV1().Pods(req.Namespace)
	podResult, err := podClient.List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	for _, pod := range podResult.Items {
		fmt.Println(pod.Name, pod.Namespace, pod.CreationTimestamp)
		list = append(list, &types.PodsListData{
			Name:              pod.Name,
			Status:            "",
			Labels:            pod.Labels["run"],
			Namespace:         pod.Namespace,
			HostIP:            pod.Status.HostIP,
			PodIP:             pod.Status.PodIP,
			StartTime:         pod.Status.StartTime.Format("2006-01-02 15:04:05"),
			RestartCount:      0,
			Image:             "",
			CreationTimestamp: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})

		logx.Info("de%v", pod)
	}

	return &types.PodsListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil

}
