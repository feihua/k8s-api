package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

type NamespaceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceDeleteLogic {
	return NamespaceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceDeleteLogic) NamespaceDelete(req types.NamespaceDeleteReq) (*types.NamespaceDeleteResp, error) {

	_ = l.svcCtx.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), req.Name, metaV1.DeleteOptions{})

	return &types.NamespaceDeleteResp{
		Code:    0,
		Message: "successful",
	}, nil
}
