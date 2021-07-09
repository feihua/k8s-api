package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceGetLogic {
	return NamespaceGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceGetLogic) NamespaceGet(req types.NamespaceGetReq) (*types.NamespaceGetResp, error) {
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	forConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	namespace, err := forConfig.CoreV1().Namespaces().Get(context.TODO(), req.Name, metaV1.GetOptions{})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.NamespaceGetResp{
		Code: 0,
		Msg:  "successful",
		Data: types.NamespaceGetData{
			Name:              namespace.Name,
			ClusterName:       namespace.ClusterName,
			Status:            string(namespace.Status.Phase),
			CreationTimestamp: namespace.CreationTimestamp.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
