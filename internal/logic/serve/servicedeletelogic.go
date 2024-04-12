package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

type ServiceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServiceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ServiceDeleteLogic {
	return ServiceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServiceDeleteLogic) ServiceDelete(req types.ServiceDeleteReq) (*types.ServiceDeleteResp, error) {

	l.svcCtx.ClientSet.CoreV1().Services(req.Namespace).Delete(context.TODO(), req.Service, metaV1.DeleteOptions{})

	return &types.ServiceDeleteResp{
		Code:    0,
		Message: "删除service: " + req.Service + "成功",
	}, nil
}
