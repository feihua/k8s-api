package jenkins

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllJobsNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllJobsNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllJobsNameLogic {
	return GetAllJobsNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllJobsNameLogic) GetAllJobsName(req types.JenkinsListReq) (*types.JenkinsListResp, error) {
	jobs, err := l.svcCtx.Jenkins.GetAllJobNames(l.ctx)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息失败,请求参数:%s,异常:%s", req.Name, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.JenkinsListItem
	for _, item := range jobs {
		list = append(list, &types.JenkinsListItem{
			Name:  item.Name,
			Url:   item.Url,
			Color: item.Color,
		})

	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询jenkins job列表信息,请求参数：%s,响应：%s", req.Name, listStr)

	return &types.JenkinsListResp{
		Code:    0,
		Message: "successful",
		Data: types.JenkinsListData{
			Items: list,
			Total: int64(len(list)),
		},
	}, nil
}
