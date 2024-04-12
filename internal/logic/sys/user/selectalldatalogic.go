package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAllDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) SelectAllDataLogic {
	return SelectAllDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAllDataLogic) SelectAllData(req types.SelectDataReq) (*types.SelectDataResp, error) {
	// todo: add your logic here and delete this line

	return &types.SelectDataResp{}, nil
}
