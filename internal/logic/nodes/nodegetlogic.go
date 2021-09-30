package logic

import (
	"context"
	"encoding/json"
	"k8s_test/internal/common/errorx"

	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type NodeGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) NodeGetLogic {
	return NodeGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeGetLogic) NodeGet(req types.NodeGetReq) (*types.NodeGetResp, error) {
	node, err := l.svcCtx.NodeGet(req.Name)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询单个node信息异常:%s", err.Error())
		return nil, errorx.NewDefaultError(err.Error())
	}

	var imagesData []*types.ImagesData
	images := node.Status.Images
	for _, image := range images {
		imagesData = append(imagesData, &types.ImagesData{
			Name: image.Names[0],
			Size: image.SizeBytes,
		})
	}

	nodeInfo := node.Status.NodeInfo
	data := types.NodesListData{
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
	}

	dataStr, _ := json.Marshal(data)
	logx.WithContext(l.ctx).Infof("查询单个node信息响应：%s", dataStr)
	return &types.NodeGetResp{
		Code: 0,
		Msg:  "successful",
		Data: data,
	}, nil

}
