package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecretUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretUpdateLogic {
	return SecretUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretUpdateLogic) SecretUpdate(req types.SecretUpdateReq) (*types.SecretUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &types.SecretUpdateResp{}, nil
}
