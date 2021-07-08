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

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetLogic {
	return GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req types.NamespaceListReq) (*types.NamespaceGetResp, error) {
	// k8s 配置
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	forConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	namespaceClient := forConfig.CoreV1().Namespaces()
	namespace, err := namespaceClient.Get(context.TODO(), req.Name, metaV1.GetOptions{})

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
