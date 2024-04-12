package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodDeleteLogic {
	return PodDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodDeleteLogic) PodDelete(req types.PodDeleteReq) (*types.PodDeleteResp, error) {
	l.svcCtx.ClientSet.CoreV1().Pods(req.Namespace).Delete(context.TODO(), req.Pod, metaV1.DeleteOptions{})

	return &types.PodDeleteResp{
		Code:    0,
		Message: "删除pod: " + req.Pod + "成功",
	}, nil
}
