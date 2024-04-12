package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NamespaceListWithLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNamespaceListWithLogic(ctx context.Context, svcCtx *svc.ServiceContext) NamespaceListWithLogic {
	return NamespaceListWithLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NamespaceListWithLogic) NamespaceListWith(req types.NamespaceListReq) (*types.NamespaceListRespWith, error) {
	client := l.svcCtx.ClientSet.CoreV1().Namespaces()
	result, err := client.List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询namespace列表信息,异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}
	now := time.Now()
	var list []*types.NamespaceListItem
	for _, namespace := range result.Items {
		list = append(list, &types.NamespaceListItem{
			Name:              namespace.Name,
			Status:            string(namespace.Status.Phase),
			Age:               now.Sub(namespace.CreationTimestamp.Time).String(),
			CreationTimestamp: namespace.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})
		//fmt.Println(namespace.Name, now.Sub(namespace.CreationTimestamp.Time))
	}

	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询namespace列表信息响应：%s", listStr)
	return &types.NamespaceListRespWith{
		Code:    0,
		Message: "successful",
		Data:    list,
	}, nil

}
