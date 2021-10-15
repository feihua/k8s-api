package logic

import (
	"context"
	pagination "k8s_test/internal/common/pagination"
	"k8s_test/internal/model"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserListLogic {
	return UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req types.ListUserReq) (*types.ListUserResp, error) {
	var list []model.User
	tx := l.svcCtx.DbClient.Limit(req.PageSize).Offset(pagination.GetPageOffset(req.Current, req.PageSize))

	if req.DeptId != 0 {
		tx.Where("dept_id = ?", req.DeptId)
	}
	tx.Find(&list)

	var count int64
	l.svcCtx.DbClient.Model(&model.User{}).Count(&count)

	items := make([]*types.ListUserItem, 0)
	for _, item := range list {
		items = append(items, &types.ListUserItem{
			Id:             item.Id,
			Name:           item.Name,
			NickName:       item.NickName,
			Avatar:         item.Avatar,
			Password:       item.Password,
			Salt:           item.Salt,
			Email:          item.Email,
			Mobile:         item.Mobile,
			DeptId:         item.DeptId,
			Status:         int64(item.Status),
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        int64(item.DelFlag),
			JobId:          int64(item.JobId),
			//RoleId:         item.RoleId,
			//RoleName:       item.RoleName,
			//JobName:        item.JobName,
			//DeptName:       item.DeptName,
		})
	}
	return &types.ListUserResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.ListUserData{
			Items: items,
			Total: count,
		},
	}, nil
}
