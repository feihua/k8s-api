// Code generated by goctl. DO NOT EDIT.
package types

type DeploymentAddReq struct {
	Name string `json:"name,optional"`
}

type DeploymentAddResp struct {
	Message string `json:"message"`
}

type DeploymentDeleteCollectionReq struct {
	Name string `json:"name,optional"`
}

type DeploymentDeleteCollectionResp struct {
	Message string `json:"message"`
}

type DeploymentDeleteReq struct {
	Name string `json:"name,optional"`
}

type DeploymentDeleteResp struct {
	Message string `json:"message"`
}

type DeploymentGetReq struct {
	Namespace  string `json:"namespace,optional"`
	Deployment string `json:"deployment"`
}

type DeploymentGetResp struct {
	Message string `json:"message"`
}

type DeploymentGetScaleReq struct {
	Name string `json:"name,optional"`
}

type DeploymentGetScaleResp struct {
	Message string `json:"message"`
}

type DeploymentListData struct {
	Name               string            `json:"name"`
	Namespace          string            `json:"namespace"`
	Labels             map[string]string `json:"labels"`
	Strategy           string            `json:"strategy"`
	Replicas           int32             `json:"replicas"`
	UpdatedReplicas    int32             `json:"updatedReplicas"`
	ReadyReplicas      int32             `json:"readyReplicas"`
	AvailableReplicas  int32             `json:"availableReplicas"`
	ObservedGeneration int64             `json:"observedGeneration"`
	CreationTimestamp  string            `json:"creationTimestamp"`
	Images             string            `json:"images"`
	ImagePullPolicy    string            `json:"imagePullPolicy"`
	Message            string            `json:"message"`
	Reason             string            `json:"reason"`
	LastUpdateTime     string            `json:"lastUpdateTime"`
	Status             string            `json:"status"`
}

type DeploymentListReq struct {
	Namespace string `json:"namespace,optional"`
}

type DeploymentListResp struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data []*DeploymentListData `json:"data"`
}

type DeploymentPatchReq struct {
	Name string `json:"name,optional"`
}

type DeploymentPatchResp struct {
	Message string `json:"message"`
}

type DeploymentUpdateReq struct {
	Name string `json:"name,optional"`
}

type DeploymentUpdateResp struct {
	Message string `json:"message"`
}

type DeploymentUpdateScaleReq struct {
	Name string `json:"name,optional"`
}

type DeploymentUpdateScaleResp struct {
	Message string `json:"message"`
}

type DeploymentUpdateStatusReq struct {
	Name string `json:"name,optional"`
}

type DeploymentUpdateStatusResp struct {
	Message string `json:"message"`
}

type DeploymentWatchReq struct {
	Name string `json:"name,optional"`
}

type DeploymentWatchResp struct {
	Message string `json:"message"`
}

type ImagesData struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type IngressAddReq struct {
	Name string `json:"name,optional"`
}

type IngressAddResp struct {
	Message string `json:"message"`
}

type IngressDeleteCollectionReq struct {
	Name string `json:"name,optional"`
}

type IngressDeleteCollectionResp struct {
	Message string `json:"message"`
}

type IngressDeleteReq struct {
	Name string `json:"name,optional"`
}

type IngressDeleteResp struct {
	Message string `json:"message"`
}

type IngressGetData struct {
	Name              string `json:"name"`
	Namespace         string `json:"namespace"`
	Host              string `json:"host"`
	ServiceName       string `json:"serviceName"`
	ServicePort       int32  `json:"servicePort"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type IngressGetReq struct {
	Namespace string `json:"namespace,optional"`
	Ingress   string `json:"ingress"`
}

type IngressGetResp struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data IngressGetData `json:"data"`
}

type IngressListData struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Host              string            `json:"host"`
	ServiceName       string            `json:"serviceName"`
	ServicePort       int32             `json:"servicePort"`
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Status            string            `json:"status"`
	Rules             string            `json:"rules"`
	Address           string            `json:"address"`
	ResourceVersion   string            `json:"resourceVersion"`
	Annotations       map[string]string `json:"annotations"`
}

type IngressListReq struct {
	Namespace string `json:"namespace,optional"`
}

type IngressListResp struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data []*IngressListData `json:"data"`
}

type IngressPatchReq struct {
	Name string `json:"name,optional"`
}

type IngressPatchResp struct {
	Message string `json:"message"`
}

type IngressUpdateReq struct {
	Name string `json:"name,optional"`
}

type IngressUpdateResp struct {
	Message string `json:"message"`
}

type IngressUpdateStatusReq struct {
	Name string `json:"name,optional"`
}

type IngressUpdateStatusResp struct {
	Message string `json:"message"`
}

type IngressWatchReq struct {
	Name string `json:"name,optional"`
}

type IngressWatchResp struct {
	Message string `json:"message"`
}

type NamespaceAddReq struct {
	Name string `json:"name"`
}

type NamespaceAddResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NamespaceDeleteReq struct {
	Name string `json:"name,optional"`
}

type NamespaceDeleteResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NamespaceGetData struct {
	Name              string `json:"name"`
	ClusterName       string `json:"clusterName"`
	Status            string `json:"status"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type NamespaceGetReq struct {
	Name string `json:"name,optional"`
}

type NamespaceGetResp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data NamespaceGetData `json:"data"`
}

type NamespaceListData struct {
	Name              string `json:"name"`
	Status            string `json:"status"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type NamespaceListReq struct {
	Name string `json:"name,optional"`
}

type NamespaceListResp struct {
	Code int                  `json:"code"`
	Msg  string               `json:"msg"`
	Data []*NamespaceListData `json:"data"`
}

type NamespacePatchReq struct {
	Name string `json:"name,optional"`
}

type NamespacePatchResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NamespaceUpdateReq struct {
	Name string `json:"name"`
}

type NamespaceUpdateResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NamespaceUpdateStatusReq struct {
	Name string `json:"name"`
}

type NamespaceUpdateStatusResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NamespaceWatchReq struct {
	Name string `json:"name,optional"`
}

type NamespaceWatchResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NodeAddReq struct {
	Name string `json:"name,optional"`
}

type NodeAddResp struct {
	Message string `json:"message"`
}

type NodeData struct {
	MachineID               string `json:"machineID"`
	SystemUUID              string `json:"systemUUID"`
	BootID                  string `json:"bootID"`
	KernelVersion           string `json:"kernelVersion"`
	OSImage                 string `json:"oSImage"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	OperatingSystem         string `json:"operatingSystem"`
	Architecture            string `json:"architecture"`
}

type NodeDeleteCollectionReq struct {
	Name string `json:"name,optional"`
}

type NodeDeleteCollectionResp struct {
	Message string `json:"message"`
}

type NodeDeleteReq struct {
	Name string `json:"name,optional"`
}

type NodeDeleteResp struct {
	Message string `json:"message"`
}

type NodeGetReq struct {
	Name string `json:"name,optional"`
}

type NodeGetResp struct {
	Message string `json:"message"`
}

type NodeListReq struct {
	Namespace string `json:"namespace,optional"`
}

type NodePatchReq struct {
	Name string `json:"name,optional"`
}

type NodePatchResp struct {
	Message string `json:"message"`
}

type NodeUpdateReq struct {
	Name string `json:"name,optional"`
}

type NodeUpdateResp struct {
	Message string `json:"message"`
}

type NodeUpdateStatusReq struct {
	Name string `json:"name,optional"`
}

type NodeUpdateStatusResp struct {
	Message string `json:"message"`
}

type NodeWatchReq struct {
	Name string `json:"name,optional"`
}

type NodeWatchResp struct {
	Message string `json:"message"`
}

type NodesListData struct {
	Name              string        `json:"name"`
	Status            string        `json:"status"`
	Memory            string        `json:"memory"`
	NodeInfo          NodeData      `json:"nodeInfo"`
	Images            []*ImagesData `json:"Images"`
	CreationTimestamp string        `json:"creationTimestamp"`
}

type NodesListReq struct {
	Name string `json:"name,optional"`
}

type NodesListResp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data []*NodesListData `json:"data"`
}

type PodAddReq struct {
	Name string `json:"name,optional"`
}

type PodAddResp struct {
	Message string `json:"message"`
}

type PodDeleteCollectionReq struct {
	Name string `json:"name,optional"`
}

type PodDeleteCollectionResp struct {
	Message string `json:"message"`
}

type PodDeleteReq struct {
	Name string `json:"name,optional"`
}

type PodDeleteResp struct {
	Message string `json:"message"`
}

type PodGetEphemeralContainersReq struct {
	Name string `json:"name,optional"`
}

type PodGetEphemeralContainersResp struct {
	Message string `json:"message"`
}

type PodGetReq struct {
	Name string `json:"name,optional"`
}

type PodGetResp struct {
	Message string `json:"message"`
}

type PodListReq struct {
	Namespace string `json:"namespace,optional"`
}

type PodPatchReq struct {
	Name string `json:"name,optional"`
}

type PodPatchResp struct {
	Message string `json:"message"`
}

type PodUpdateEphemeralContainersReq struct {
	Name string `json:"name,optional"`
}

type PodUpdateEphemeralContainersResp struct {
	Message string `json:"message"`
}

type PodUpdateReq struct {
	Name string `json:"name,optional"`
}

type PodUpdateResp struct {
	Message string `json:"message"`
}

type PodUpdateStatusReq struct {
	Name string `json:"name,optional"`
}

type PodUpdateStatusResp struct {
	Message string `json:"message"`
}

type PodWatchReq struct {
	Name string `json:"name,optional"`
}

type PodWatchResp struct {
	Message string `json:"message"`
}

type PodsListData struct {
	Name               string            `json:"name"`
	Status             string            `json:"status"`
	Labels             map[string]string `json:"labels"`
	NodeSelector       map[string]string `json:"nodeSelector"`
	Namespace          string            `json:"namespace"`
	HostIP             string            `json:"hostIP"`
	PodIP              string            `json:"podIP"`
	StartTime          string            `json:"startTime"`
	RestartCount       int32             `json:"restartCount"`
	Image              string            `json:"image"`
	CreationTimestamp  string            `json:"creationTimestamp"`
	RestartPolicy      string            `json:"restartPolicy"`
	ServiceAccountName string            `json:"serviceAccountName"`
	NodeName           string            `json:"nodeName"`
	HostNetwork        bool              `json:"hostNetwork"`
	ImagePullSecrets   string            `json:"imagePullSecrets"`
	Hostname           string            `json:"hostname"`
}

type PodsListReq struct {
	Namespace string `json:"namespace,optional"`
}

type PodsListResp struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data []*PodsListData `json:"data"`
}

type SecretAddReq struct {
	Name string `json:"name"`
}

type SecretAddResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SecretDeleteReq struct {
	Name string `json:"name,optional"`
}

type SecretDeleteResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SecretGetData struct {
	Name              string `json:"name"`
	ClusterName       string `json:"clusterName"`
	Status            string `json:"status"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type SecretGetReq struct {
	Name string `json:"name,optional"`
}

type SecretGetResp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data SecretGetData `json:"data"`
}

type SecretListData struct {
	Name              string `json:"name"`
	NameSpace         string `json:"nameSpace"`
	Type              string `json:"type"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type SecretListReq struct {
	Namespace string `json:"namespace,optional"`
}

type SecretListResp struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data []*SecretListData `json:"data"`
}

type SecretPatchReq struct {
	Name string `json:"name,optional"`
}

type SecretPatchResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SecretUpdateReq struct {
	Name string `json:"name"`
}

type SecretUpdateResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SecretUpdateStatusReq struct {
	Name string `json:"name"`
}

type SecretUpdateStatusResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SecretWatchReq struct {
	Name string `json:"name,optional"`
}

type SecretWatchResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ServiceAddReq struct {
	Name string `json:"name,optional"`
}

type ServiceAddResp struct {
	Message string `json:"message"`
}

type ServiceDeleteCollectionReq struct {
	Name string `json:"name,optional"`
}

type ServiceDeleteCollectionResp struct {
	Message string `json:"message"`
}

type ServiceDeleteReq struct {
	Name string `json:"name,optional"`
}

type ServiceDeleteResp struct {
	Message string `json:"message"`
}

type ServiceGetReq struct {
	Name string `json:"name,optional"`
}

type ServiceGetResp struct {
	Message string `json:"message"`
}

type ServiceListData struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	Selector          map[string]string `json:"selector"`
	Type              string            `json:"type"`
	ClusterIP         string            `json:"clusterIP"`
	Protocol          string            `json:"protocol"`
	Ports             int32             `json:"ports"`
	TargetPort        int32             `json:"targetPort"`
	NodePort          int32             `json:"nodePort"`
	CreationTimestamp string            `json:"creationTimestamp"`
}

type ServiceListReq struct {
	Namespace string `json:"namespace,optional"`
}

type ServiceListResp struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data []*ServiceListData `json:"data"`
}

type ServicePatchReq struct {
	Name string `json:"name,optional"`
}

type ServicePatchResp struct {
	Message string `json:"message"`
}

type ServiceUpdateReq struct {
	Name string `json:"name,optional"`
}

type ServiceUpdateResp struct {
	Message string `json:"message"`
}

type ServiceUpdateStatusReq struct {
	Name string `json:"name,optional"`
}

type ServiceUpdateStatusResp struct {
	Message string `json:"message"`
}

type ServiceWatchReq struct {
	Name string `json:"name,optional"`
}

type ServiceWatchResp struct {
	Message string `json:"message"`
}

type StatefulSetAddReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetAddResp struct {
	Message string `json:"message"`
}

type StatefulSetDeleteCollectionReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetDeleteCollectionResp struct {
	Message string `json:"message"`
}

type StatefulSetDeleteReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetDeleteResp struct {
	Message string `json:"message"`
}

type StatefulSetGetReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetGetResp struct {
	Message string `json:"message"`
}

type StatefulSetListData struct {
	Name              string `json:"name"`
	Namespace         string `json:"namespace"`
	Labels            string `json:"labels"`
	Selector          string `json:"selector"`
	Type              string `json:"type"`
	ClusterIP         string `json:"clusterIP"`
	Protocol          string `json:"protocol"`
	Ports             int32  `json:"ports"`
	TargetPort        int32  `json:"targetPort"`
	NodePort          int32  `json:"nodePort"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type StatefulSetListReq struct {
	Namespace string `json:"namespace,optional"`
}

type StatefulSetListResp struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data []*StatefulSetListData `json:"data"`
}

type StatefulSetPatchReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetPatchResp struct {
	Message string `json:"message"`
}

type StatefulSetUpdateReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetUpdateResp struct {
	Message string `json:"message"`
}

type StatefulSetUpdateStatusReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetUpdateStatusResp struct {
	Message string `json:"message"`
}

type StatefulSetWatchReq struct {
	Name string `json:"name,optional"`
}

type StatefulSetWatchResp struct {
	Message string `json:"message"`
}
