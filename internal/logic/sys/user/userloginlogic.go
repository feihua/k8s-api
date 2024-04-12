package logic

import (
	"context"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &types.LoginResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.LoginData{
			Status:           "",
			CurrentAuthority: "",
			Id:               1,
			UserName:         "liufeihua",
			AccessToken:      "123456",
			AccessExpire:     0,
			RefreshAfter:     0,
		},
	}, nil
}
