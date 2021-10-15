package logic

import (
	"context"
	"k8s_test/internal/common/pagination"
	"k8s_test/internal/model"

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
	var list []model.Dict
	l.svcCtx.DbClient.Limit(req.PageSize).Offset(pagination.GetPageOffset(req.Current, req.PageSize)).Find(&list)

	var count int64
	l.svcCtx.DbClient.Model(&model.Dict{}).Count(&count)

	items := make([]*types.ListDictItem, 0)
	for _, item := range list {
		items = append(items, &types.ListDictItem{
			Id:             item.Id,
			Value:          item.Value,
			Label:          item.Label,
			Type:           item.Type,
			Description:    item.Description,
			Sort:           item.Sort,
			Remarks:        item.Remarks,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        item.DelFlag,
		})
	}
	return &types.ListDictResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.ListDictData{
			Items: items,
			Total: count,
		},
	}, nil

}
