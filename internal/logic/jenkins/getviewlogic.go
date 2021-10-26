package jenkins

import (
	"context"
	"encoding/json"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetViewLogic {
	return GetViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetViewLogic) GetView(req types.GetViewReq) (*types.GetViewResp, error) {
	view, err := l.svcCtx.Jenkins.GetView(l.ctx, req.Name)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息失败,请求参数:%s,异常:%s", req.Name, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var result []*types.JenkinsListItem
	raw, _ := json.Marshal(view.Raw.Jobs)
	json.Unmarshal(raw, &result)

	return &types.GetViewResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.GetViewData{
			Items: result,
			Total: int64(len(result)),
		},
	}, nil

}
