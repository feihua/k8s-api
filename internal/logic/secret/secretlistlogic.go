package logic

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/common/errorx"
	"log"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SecretListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var TypeSelect = map[string]string{
	"Opaque":                              "自定义类型",
	"kubernetes.io/service-account-token": "服务账号令牌",
	"kubernetes.io/dockercfg":             "docker配置",
	"kubernetes.io/dockerconfigjson":      "docker配置(JSON)",
	"kubernetes.io/basic-auth":            "Basic认证凭据",
	"kubernetes.io/ssh-auth":              "SSH凭据",
	"kubernetes.io/tls":                   "TLS凭据",
	"bootstrap.kubernetes.io/token":       "启动引导令牌数据",
}

func NewSecretListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretListLogic {
	return SecretListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretListLogic) SecretList(req types.SecretListReq) (*types.SecretListResp, error) {
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

	// 5. service 列表
	fmt.Println("services:")
	serviceClient := clientSet.CoreV1().Secrets(req.Namespace)
	serviceResult, err := serviceClient.List(context.TODO(), metaV1.ListOptions{})

	var list []*types.SecretListData
	if err != nil {
		log.Println(err)
	} else {
		for _, item := range serviceResult.Items {
			list = append(list, &types.SecretListData{
				Name:              item.Name,
				NameSpace:         item.Namespace,
				Type:              TypeSelect[string(item.Type)],
				CreationTimestamp: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			})
		}
	}

	return &types.SecretListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil

}
