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

type GetAllJobsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllJobsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllJobsLogic {
	return GetAllJobsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllJobsLogic) GetAllJobs() (*types.GetAllJobsResp, error) {
	jobs, err := l.svcCtx.Jenkins.GetAllJobs(l.ctx)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job列表信息异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var result []*types.Result
	for _, item := range jobs {
		var i types.Result
		raw, _ := json.Marshal(item.Raw)
		json.Unmarshal(raw, &i)
		result = append(result, &i)
		fmt.Println(string(raw))

	}

	return &types.GetAllJobsResp{
		Code:    0,
		Message: "successful",
		Data: types.GetAllJobsData{
			Items: result,
			Total: int64(len(result)),
		},
	}, nil

}
