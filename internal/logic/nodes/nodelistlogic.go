package logic

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
	kubeConfig := "etc/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	var listData []*types.NodesListData

	nodes, err := clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

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
		listData = append(listData, &types.NodesListData{
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

		fmt.Println(node.Name)
		fmt.Println(node.CreationTimestamp) //加入集群时间
		fmt.Println(nodeInfo)
		fmt.Println(node.Status.Conditions[len(node.Status.Conditions)-1].Type)
		fmt.Println(node.Status.Allocatable.Memory().String())
	}

	return &types.NodesListResp{
		Code: 0,
		Msg:  "successful",
		Data: listData,
	}, nil
}
