package logic

import (
	"context"
	"fmt"
	apiV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NamespaceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceCreateLogic {
	return NamespaceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceCreateLogic) NamespaceCreate(req types.NamespaceAddReq) (*types.NamespaceAddResp, error) {
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

	namespaceClient := clientSet.CoreV1().Namespaces()
	namespace := &apiV1.Namespace{
		ObjectMeta: metaV1.ObjectMeta{
			Name: req.Name,
		},
		Status: apiV1.NamespaceStatus{
			Phase: apiV1.NamespaceActive,
		},
	}
	result, err := namespaceClient.Create(context.TODO(), namespace, metaV1.CreateOptions{})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println(result)
	return &types.NamespaceAddResp{
		Code: 0,
		Msg:  "successful",
	}, nil
}