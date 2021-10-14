package logic

import (
	"context"
	"encoding/json"
	apiV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	namespace := &apiV1.Namespace{
		ObjectMeta: metaV1.ObjectMeta{
			Name: req.Name,
		},
		Status: apiV1.NamespaceStatus{
			Phase: apiV1.NamespaceActive,
		},
	}

	client := l.svcCtx.ClientSet.CoreV1().Namespaces()
	result, err := client.Create(context.TODO(), namespace, metaV1.CreateOptions{})

	if err != nil {
		resultStr, _ := json.Marshal(result)
		logx.WithContext(l.ctx).Errorf("创建namespace信息失败,请求参数:%s,异常:%s", resultStr, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.NamespaceAddResp{
		Code:    0,
		Message: "successful",
	}, nil
}
