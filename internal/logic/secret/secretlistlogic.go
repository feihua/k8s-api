package logic

import (
	"context"
	"encoding/json"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"
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
	result, err := l.svcCtx.ClientSet.CoreV1().Secrets(req.Namespace).List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询secret列表信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.SecretListData
	for _, item := range result.Items {
		list = append(list, &types.SecretListData{
			Name:              item.Name,
			NameSpace:         item.Namespace,
			Type:              TypeSelect[string(item.Type)],
			CreationTimestamp: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询secret列表信息,请求参数：%s,响应：%s", req.Namespace, listStr)
	return &types.SecretListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil

}
