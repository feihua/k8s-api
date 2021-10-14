package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptDeleteLogic {
	return DeptDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptDeleteLogic) DeptDelete(req types.DeleteDeptReq) (*types.DeleteDeptResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteDeptResp{}, nil
}
