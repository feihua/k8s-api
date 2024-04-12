package logic

import (
	"context"
	"github.com/feihua/k8s-api/internal/common/pagination"
	"github.com/feihua/k8s-api/internal/model"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogListLogic {
	return LoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogListLogic) LoginLogList(req types.ListLoginLogReq) (*types.ListLoginLogResp, error) {
	var list []model.LoginLog
	l.svcCtx.DbClient.Limit(req.PageSize).Offset(pagination.GetPageOffset(req.Current, req.PageSize)).Find(&list)

	var count int64
	l.svcCtx.DbClient.Model(&model.LoginLog{}).Count(&count)

	items := make([]*types.ListLoginLogItem, 0)
	for _, item := range list {
		items = append(items, &types.ListLoginLogItem{
			Id:             item.Id,
			UserName:       item.UserName,
			Status:         item.Status,
			Ip:             item.Ip,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &types.ListLoginLogResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.ListLoginLogData{
			Items: items,
			Total: count,
		},
	}, nil

}
