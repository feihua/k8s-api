package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/k8s-api/internal/common/errorx"

	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeListLogic {
	return NodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeListLogic) NodeList(req types.NodesListReq) (*types.NodesListResp, error) {
	nodes, err := l.svcCtx.NodeList()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询node列表信息异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.NodesListItem
	for _, node := range nodes.Items {

		nodeInfo := node.Status.NodeInfo
		list = append(list, &types.NodesListItem{
			Name:                    node.Name,
			Status:                  string(node.Status.Conditions[len(node.Status.Conditions)-1].Type),
			OSImage:                 nodeInfo.OSImage,
			ContainerRuntimeVersion: nodeInfo.ContainerRuntimeVersion,
			KubeletVersion:          nodeInfo.KubeletVersion,
			Architecture:            nodeInfo.Architecture,
			CreationTimestamp:       node.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})

	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询node列表信息响应：%s", listStr)
	return &types.NodesListResp{
		Code:    0,
		Message: "ok",
		Type:    "success",
		Data: types.NodesListData{
			Items: list,
			Total: int64(len(list)),
		},
	}, nil
}
