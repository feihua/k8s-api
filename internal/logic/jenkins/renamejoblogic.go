package jenkins

import (
	"context"
	"fmt"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RenameJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRenameJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) RenameJobLogic {
	return RenameJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RenameJobLogic) RenameJob(req types.RenameJobReq) (*types.RenameJobResp, error) {
	job := l.svcCtx.Jenkins.RenameJob(l.ctx, req.OldName, req.NewName)

	fmt.Println(job)
	return &types.RenameJobResp{
		Code:    0,
		Message: "successful",
	}, nil

}
