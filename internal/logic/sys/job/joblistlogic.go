package logic

import (
	"context"
	"github.com/feihua/k8s-api/internal/common/pagination"
	"github.com/feihua/k8s-api/internal/model"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) JobListLogic {
	return JobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JobListLogic) JobList(req types.ListJobReq) (*types.ListJobResp, error) {
	var list []model.Job
	l.svcCtx.DbClient.Limit(req.PageSize).Offset(pagination.GetPageOffset(req.Current, req.PageSize)).Find(&list)

	var count int64
	l.svcCtx.DbClient.Model(&model.Job{}).Count(&count)

	items := make([]*types.ListJobItem, 0)
	for _, item := range list {
		items = append(items, &types.ListJobItem{
			Id:             item.Id,
			JobName:        item.JobName,
			OrderNum:       item.OrderNum,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        item.DelFlag,
			Remarks:        item.Remarks,
		})
	}
	return &types.ListJobResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.ListJobData{
			Items: items,
			Total: count,
		},
	}, nil

}
