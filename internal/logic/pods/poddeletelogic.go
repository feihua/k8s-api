package logic

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
		Code: 0,
		Msg:  "删除pod: " + req.Pod + "成功",
	}, nil
}
