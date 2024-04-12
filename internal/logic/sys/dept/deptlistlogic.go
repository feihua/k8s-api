package logic

import (
	"context"
	"github.com/feihua/k8s-api/internal/model"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeptListLogic {
	return DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList(req types.ListDeptReq) (*types.ListDeptResp, error) {
	var list []model.Dept
	l.svcCtx.DbClient.Find(&list)

	items := make([]*types.ListDeptData, 0)
	for _, item := range list {
		items = append(items, &types.ListDeptData{
			Id:             item.Id,
			Name:           item.Name,
			ParentId:       item.ParentId,
			OrderNum:       item.OrderNum,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        item.DelFlag,
		})
	}
	return &types.ListDeptResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data:    items,
	}, nil

}
