package jenkins

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/feihua/k8s-api/internal/common/errorx"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetInfoLogic {
	return GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo() (*types.JenkinsInfoResp, error) {
	info, err := l.svcCtx.Jenkins.Info(l.ctx)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins info信息异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	bytes, err := json.Marshal(info)
	fmt.Println(string(bytes))

	return &types.JenkinsInfoResp{
		Code:    0,
		Message: "successful",
	}, nil
}
