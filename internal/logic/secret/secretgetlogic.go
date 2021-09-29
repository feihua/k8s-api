package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SecretGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretGetLogic {
	return SecretGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretGetLogic) SecretGet(req types.SecretGetReq) (*types.SecretGetResp, error) {
	// todo: add your logic here and delete this line

	return &types.SecretGetResp{}, nil
}
