package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
