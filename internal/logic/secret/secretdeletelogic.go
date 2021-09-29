package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SecretDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretDeleteLogic {
	return SecretDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretDeleteLogic) SecretDelete(req types.SecretDeleteReq) (*types.SecretDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &types.SecretDeleteResp{}, nil
}
