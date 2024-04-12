package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"
	apiV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecretCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretCreateLogic {
	return SecretCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretCreateLogic) SecretCreate(req types.SecretAddReq) (*types.SecretAddResp, error) {
	data := make(map[string][]byte)
	data["user"] = []byte(req.User)
	data["password"] = []byte(req.Password)

	secret := &apiV1.Secret{
		TypeMeta: metaV1.TypeMeta{
			Kind: "Secret",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
		},
		Data: data,
	}

	result, err := l.svcCtx.ClientSet.CoreV1().Secrets(req.Namespace).Create(context.TODO(), secret, metaV1.CreateOptions{})

	if err != nil {
		resultStr, _ := json.Marshal(result)
		logx.WithContext(l.ctx).Errorf("创建secret信息失败,请求参数:%s,异常:%s", resultStr, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.SecretAddResp{
		Code:    0,
		Message: "successful",
	}, nil
}
