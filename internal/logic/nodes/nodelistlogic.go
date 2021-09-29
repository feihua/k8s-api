package logic

import (
	"context"
	"encoding/json"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	nodes, err := l.svcCtx.ClientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询node列表信息异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var list []*types.NodesListData
	for _, node := range nodes.Items {

		var imagesData []*types.ImagesData
		images := node.Status.Images
		for _, image := range images {
			imagesData = append(imagesData, &types.ImagesData{
				Name: image.Names[0],
				Size: image.SizeBytes,
			})
		}

		nodeInfo := node.Status.NodeInfo
		list = append(list, &types.NodesListData{
			Name:   node.Name,
			Status: string(node.Status.Conditions[len(node.Status.Conditions)-1].Type),
			Memory: node.Status.Allocatable.Memory().String(),
			NodeInfo: types.NodeData{
				MachineID:               nodeInfo.MachineID,
				SystemUUID:              nodeInfo.SystemUUID,
				BootID:                  nodeInfo.BootID,
				KernelVersion:           nodeInfo.KernelVersion,
				OSImage:                 nodeInfo.OSImage,
				ContainerRuntimeVersion: nodeInfo.ContainerRuntimeVersion,
				KubeletVersion:          nodeInfo.KubeletVersion,
				KubeProxyVersion:        nodeInfo.KubeProxyVersion,
				OperatingSystem:         nodeInfo.OperatingSystem,
				Architecture:            nodeInfo.Architecture,
			},
			Images:            imagesData,
			CreationTimestamp: node.CreationTimestamp.Format("2006-01-02 15:04:05"),
		})

	}
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询node列表信息响应：%s", listStr)
	return &types.NodesListResp{
		Code: 0,
		Msg:  "successful",
		Data: list,
	}, nil
}
