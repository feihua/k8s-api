package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SecretWatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretWatchLogic {
	return SecretWatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretWatchLogic) SecretWatch(req types.SecretWatchReq) (*types.SecretWatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.SecretWatchResp{}, nil
}
