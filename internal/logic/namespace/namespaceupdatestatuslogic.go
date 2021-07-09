package logic

import (
	"context"
	apiV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceUpdateStatusLogic {
	return NamespaceUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceUpdateStatusLogic) NamespaceUpdateStatus(req types.NamespaceUpdateStatusReq) (*types.NamespaceUpdateStatusResp, error) {
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

	namespace := &apiV1.Namespace{
		ObjectMeta: metaV1.ObjectMeta{
			Name: req.Name,
		},
		Status: apiV1.NamespaceStatus{
			Phase: apiV1.NamespaceActive,
		},
	}

	_, err = namespaceClient.UpdateStatus(context.TODO(), namespace, metaV1.UpdateOptions{})

	return &types.NamespaceUpdateStatusResp{}, nil
}
