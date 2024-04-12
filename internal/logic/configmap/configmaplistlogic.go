package logic

import (
	"context"
	"encoding/json"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigMapListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigMapListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConfigMapListLogic {
	return ConfigMapListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigMapListLogic) ConfigMapList(req types.ConfigMapListReq) (*types.ConfigMapListResp, error) {
	client := l.svcCtx.ClientSet.CoreV1().ConfigMaps(req.Namespace)
	result, err := client.List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询configmap列表信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.ConfigMapListItem
	for _, item := range result.Items {
		list = append(list, &types.ConfigMapListItem{
			Name:              item.Name,
			NameSpace:         item.Namespace,
			Labels:            item.Labels,
			Annotations:       item.Annotations,
			CreationTimestamp: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
			Data:              item.Data,
		})
	}

	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询configmap列表信息,请求参数：%s,响应：%s", req.Namespace, listStr)
	return &types.ConfigMapListResp{
		Code:    0,
		Message: "successful",
		Data: types.ConfigMapListData{
			Items: list,
			Total: int64(len(list)),
		},
	}, nil

}
