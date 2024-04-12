package logic

import (
	"context"
	"k8s_test/internal/model"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuListLogic {
	return MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList(req types.ListMenuReq) (*types.ListMenuResp, error) {
	var list []model.Menu
	l.svcCtx.DbClient.Find(&list)

	items := make([]*types.ListtMenuData, 0)
	for _, item := range list {
		items = append(items, &types.ListtMenuData{
			Id:             item.Id,
			Name:           item.Name,
			ParentId:       item.ParentId,
			Url:            item.Url,
			Perms:          item.Perms,
			Type:           item.Type,
			Icon:           item.Icon,
			OrderNum:       item.OrderNum,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        item.DelFlag,
		})
	}
	return &types.ListMenuResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data:    items,
	}, nil

}
