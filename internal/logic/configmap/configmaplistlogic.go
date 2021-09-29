package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"log"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ConfigMapListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapListLogic {
	return ConfigMapListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapListLogic) ConfigMapList(req types.ConfigMapListReq) (*types.ConfigMapListResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	serviceClient := clientSet.CoreV1().ConfigMaps(req.Namespace)
	serviceResult, err := serviceClient.List(context.TODO(), metaV1.ListOptions{})

	var list []*types.ConfigMapListData
	if err != nil {
		log.Println(err)
	} else {
		for _, item := range serviceResult.Items {
			list = append(list, &types.ConfigMapListData{
				Name:              item.Name,
				NameSpace:         item.Namespace,
				Labels:            item.Labels,
				Annotations:       item.Annotations,
				CreationTimestamp: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
				Data:              item.Data,
			})
		}
	}

	return &types.ConfigMapListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil

}
