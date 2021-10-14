package logic

import (
	"context"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &types.UserInfoResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.UserInfoData{
			Id:       "1",
			Username: "koobe",
			RealName: "liufeihua",
			Avatar:   "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640",
			Desc:     "测试",
		},
	}, nil
}
