package logic

import (
	"context"
	"fmt"
	"github.com/feihua/k8s-api/internal/common/errorx"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"

	"github.com/zeromicro/go-zero/core/logx"
)

type PodLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPodLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) PodLogLogic {
	return PodLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PodLogLogic) PodLog(req types.PodLogReq) (*types.PodLogResp, error) {

	var (
		result rest.Result
		data   []byte
		err    error
	)

	logs := l.svcCtx.ClientSet.CoreV1().Pods(req.Namespace).GetLogs(req.PodName, &coreV1.PodLogOptions{Container: req.ContainerName, TailLines: &req.TailLines})

	if result = logs.Do(context.TODO()); result.Error() != nil {
		logx.WithContext(l.ctx).Errorf("获取pod控制台日志输出信息失败,请求参数:%s,异常:%s", req.Namespace, result.Error())
		return nil, errorx.NewDefaultError(result.Error().Error())
	}

	if data, err = result.Raw(); err != nil {
		logx.WithContext(l.ctx).Errorf("获取pod控制台日志输出信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println("容器输出:", string(data))
	return &types.PodLogResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data:    string(data),
	}, nil
}
