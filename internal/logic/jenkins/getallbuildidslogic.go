package jenkins

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllBuildIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllBuildIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllBuildIdsLogic {
	return GetAllBuildIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllBuildIdsLogic) GetAllBuildIds(req types.GetAllBuildIdsReq) (*types.GetAllBuildIdsResp, error) {
	build, err := l.svcCtx.Jenkins.GetAllBuildIds(l.ctx, req.JobName)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询job all build列表信息失败,请求参数:%s,异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var data []*types.GetAllBuildIdsItem
	bytes, _ := json.Marshal(build)
	json.Unmarshal(bytes, &data)

	return &types.GetAllBuildIdsResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.GetAllBuildIdsData{
			Items: data,
			Total: int64(len(data)),
		},
	}, nil

}
