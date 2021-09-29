package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SecretCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretCreateLogic {
	return SecretCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretCreateLogic) SecretCreate(req types.SecretAddReq) (*types.SecretAddResp, error) {
	// todo: add your logic here and delete this line

	return &types.SecretAddResp{}, nil
}
