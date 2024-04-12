package jenkins

import (
	"context"
	"fmt"
	"github.com/feihua/k8s-api/internal/common/errorx"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) CopyJobLogic {
	return CopyJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyJobLogic) CopyJob(req types.CopyJobReq) (*types.CopyJobResp, error) {
	job, err := l.svcCtx.Jenkins.CopyJob(l.ctx, req.CopyFrom, req.NewName)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("复制jenkins job失败,请求参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	fmt.Println(job)
	return &types.CopyJobResp{
		Code:    0,
		Message: "successful",
	}, nil

}
