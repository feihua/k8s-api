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

type GetJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetJobLogic {
	return GetJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJobLogic) GetJob(req types.GetJobReq) (*types.GetJobResp, error) {
	jobs, err := l.svcCtx.Jenkins.GetJob(l.ctx, req.Id)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询jenkins job信息失败,请求参数:%s,异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var result types.Result
	raw, _ := json.Marshal(jobs.Raw)
	json.Unmarshal(raw, &result)

	fmt.Println(jobs)
	return &types.GetJobResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data:    result,
	}, nil

}
