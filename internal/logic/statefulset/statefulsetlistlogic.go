package logic

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

type StatefulSetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatefulSetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) StatefulSetListLogic {
	return StatefulSetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatefulSetListLogic) StatefulSetList(req types.StatefulSetListReq) (*types.StatefulSetListResp, error) {
	client := l.svcCtx.ClientSet.AppsV1().StatefulSets(req.Namespace)
	result, err := client.List(context.TODO(), metaV1.ListOptions{})

	var list []*types.StatefulSetListItem
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询statefulSets列表信息失败,请求参数:%s,异常:%s", req.Namespace, err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	for _, item := range result.Items {
		labelsStr, _ := json.Marshal(item.Labels)
		list = append(list, &types.StatefulSetListItem{
			Name:              item.Name,
			NameSpace:         item.Namespace,
			ClusterName:       item.ClusterName,
			Labels:            item.Labels,
			Annotations:       item.Annotations,
			LabelsStr:         string(labelsStr),
			CreationTimestamp: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询statefulSets列表信息,请求参数：%s,响应：%s", req.Namespace, listStr)
	return &types.StatefulSetListResp{
		Code:    0,
		Message: "successful",
		Data: types.StatefulSetListData{
			Items: list,
			Total: int64(len(list)),
		},
	}, nil

}
