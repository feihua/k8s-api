package jenkins

import (
	"context"
	"encoding/json"
	"fmt"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllViewsWithLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllViewsWithLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllViewsWithLogic {
	return GetAllViewsWithLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllViewsWithLogic) GetAllViewsWith() (*types.GetAllViewsRespWith, error) {
	views, err := l.svcCtx.Jenkins.GetAllViews(l.ctx)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}
	var result []*types.AllViewsListItem
	for _, item := range views {
		var i types.AllViewsListItem
		raw, _ := json.Marshal(item.Raw)
		json.Unmarshal(raw, &i)
		result = append(result, &i)
		fmt.Println(string(raw))

	}

	return &types.GetAllViewsRespWith{
		Code:    0,
		Message: "successful",
		Data:    result,
	}, nil
}
