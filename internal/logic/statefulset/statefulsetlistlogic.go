package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
	"log"

	"github.com/tal-tech/go-zero/core/logx"
)

type StatefulSetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetListLogic {
	return StatefulSetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetListLogic) StatefulSetList(req types.StatefulSetListReq) (*types.StatefulSetListResp, error) {
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	serviceClient := clientSet.AppsV1().StatefulSets(req.Namespace)
	serviceResult, err := serviceClient.List(context.TODO(), metaV1.ListOptions{})

	var list []*types.StatefulSetListData
	if err != nil {
		log.Println(err)
	} else {
		for _, item := range serviceResult.Items {
			list = append(list, &types.StatefulSetListData{
				Name:              item.Name,
				NameSpace:         item.Namespace,
				ClusterName:       item.ClusterName,
				Labels:            item.Labels,
				Annotations:       item.Annotations,
				CreationTimestamp: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			})
		}
	}

	return &types.StatefulSetListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil

}
