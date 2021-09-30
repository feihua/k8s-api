package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IngressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIngressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) IngressDeleteLogic {
	return IngressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IngressDeleteLogic) IngressDelete(req types.IngressDeleteReq) (*types.IngressDeleteResp, error) {
	client := l.svcCtx.ClientSet.ExtensionsV1beta1().Ingresses(req.Namespace)

	err := client.Delete(context.TODO(), req.Ingress, metaV1.DeleteOptions{})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.IngressDeleteResp{
		Code: 0,
		Msg:  "删除ingress: " + req.Ingress + "成功",
	}, nil
}
