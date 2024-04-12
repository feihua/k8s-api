package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecretUpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretUpdateStatusLogic {
	return SecretUpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretUpdateStatusLogic) SecretUpdateStatus(req types.SecretUpdateStatusReq) (*types.SecretUpdateStatusResp, error) {
	// todo: add your logic here and delete this line

	return &types.SecretUpdateStatusResp{}, nil
}
