package logic

import (
	"context"
	"k8s_test/internal/common/pagination"
	"k8s_test/internal/model"
	"strconv"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleListLogic {
	return RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req types.ListRoleReq) (*types.ListRoleResp, error) {
	var list []model.Role
	l.svcCtx.DbClient.Limit(req.PageSize).Offset(pagination.GetPageOffset(req.Current, req.PageSize)).Find(&list)

	var count int64
	l.svcCtx.DbClient.Model(&model.Role{}).Count(&count)

	items := make([]*types.ListRoleItem, 0)
	for _, item := range list {
		items = append(items, &types.ListRoleItem{
			Id:         item.Id,
			RoleName:   item.Name,
			RoleValue:  item.Value,
			OrderNo:    item.Sort,
			Remark:     item.Remark,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
			Status:     strconv.FormatInt(int64(item.Status), 10),
		})
	}
	return &types.ListRoleResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.ListRoleData{
			Items: items,
			Total: count,
		},
	}, nil
}
