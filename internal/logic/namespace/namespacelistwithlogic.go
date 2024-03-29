package logic

import (
	"context"
	"encoding/json"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"
	"time"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
