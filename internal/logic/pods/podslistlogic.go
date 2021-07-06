package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

type PodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodsListLogic {
	return PodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodsListLogic) PodsList(req types.PodsListReq) (*types.PodsListResp, error) {

	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	//获取POD
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	var listData []*types.PodsListData
	fmt.Println("pod:")
	for _, pod := range pods.Items {
		listData = append(listData, &types.PodsListData{
			Name:              pod.Name,
			Status:            string(pod.Status.Phase),
			Labels:            pod.Labels["run"],
			Namespace:         pod.Namespace,
			HostIP:            pod.Status.HostIP,
			PodIP:             pod.Status.PodIP,
			StartTime:         pod.Status.StartTime.Format("2006-01-02 15:04:05"),
			RestartCount:      pod.Status.ContainerStatuses[0].RestartCount,
			Image:             pod.Status.ContainerStatuses[0].Image,
			CreationTimestamp: pod.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})

		fmt.Println(pod.Name)
		fmt.Println(pod.CreationTimestamp)
		fmt.Println(pod.Labels)
		fmt.Println(pod.Namespace)
		fmt.Println(pod.Status.HostIP)
		fmt.Println(pod.Status.PodIP)
		fmt.Println(pod.Status.StartTime)
		fmt.Println(pod.Status.Phase)
		fmt.Println(pod.Status.ContainerStatuses[0].RestartCount) //重启次数
		fmt.Println(pod.Status.ContainerStatuses[0].Image)        //获取重启时间
	}

	return &types.PodsListResp{
		Code: 0,
		Msg:  "successful",
		Data: listData,
	}, nil

}
