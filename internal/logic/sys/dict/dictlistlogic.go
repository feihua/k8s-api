package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DictListLogic {
	return DictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictListLogic) DictList(req types.ListDictReq) (*types.ListDictResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListDictResp{}, nil
}
