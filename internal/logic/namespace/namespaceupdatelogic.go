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

type NamespaceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceUpdateLogic {
	return NamespaceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceUpdateLogic) NamespaceUpdate(req types.NamespaceUpdateReq) (*types.NamespaceUpdateResp, error) {
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

	_, err = namespaceClient.Update(context.TODO(), namespace, metaV1.UpdateOptions{})

	return &types.NamespaceUpdateResp{}, nil
}
