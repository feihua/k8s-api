package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SecretPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSecretPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SecretPatchLogic {
	return SecretPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SecretPatchLogic) SecretPatch(req types.SecretWatchReq) (*types.SecretPatchResp, error) {
	// todo: add your logic here and delete this line

	return &types.SecretPatchResp{}, nil
}
