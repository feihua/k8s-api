// Code generated by goctl. DO NOT EDIT.
package types

type DeploymentAddReq struct {
	Name string `json:"name,options=you|me"`
}

type DeploymentAddResp struct {
	Message string `json:"message"`
}

type DeploymentDeleteReq struct {
	Name string `json:"name,options=you|me"`
}

type DeploymentDeleteResp struct {
	Message string `json:"message"`
}

type DeploymentListData struct {
	Name              string `json:"name"`
	Namespace         string `json:"namespace"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type DeploymentListReq struct {
	Namespace string `json:"namespace"`
}

type DeploymentListResp struct {
	ListData []*DeploymentListData `json:"listData"`
}

type DeploymentUpdateReq struct {
	Name string `json:"name,options=you|me"`
}

type DeploymentUpdateResp struct {
	Message string `json:"message"`
}

type IngressAddReq struct {
	Name string `json:"name,options=you|me"`
}

type IngressAddResp struct {
	Message string `json:"message"`
}

type IngressDeleteReq struct {
	Name string `json:"name,options=you|me"`
}

type IngressDeleteResp struct {
	Message string `json:"message"`
}

type IngressListReq struct {
	Name string `json:"name,options=you|me"`
}

type IngressListResp struct {
	Message string `json:"message"`
}

type NamespaceAddReq struct {
	Name string `json:"name,options=you|me"`
}

type NamespaceAddResp struct {
	Message string `json:"message"`
}

type NamespaceDeleteReq struct {
	Name string `json:"name,options=you|me"`
}

type NamespaceDeleteResp struct {
	Message string `json:"message"`
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

type ServiceAddReq struct {
	Name string `json:"name,options=you|me"`
}

type ServiceAddResp struct {
	Message string `json:"message"`
}

type ServiceDeleteReq struct {
	Name string `json:"name,options=you|me"`
}

type ServiceDeleteResp struct {
	Message string `json:"message"`
}

type ServiceListData struct {
	Name              string `json:"name"`
	Namespace         string `json:"namespace"`
	Labels            string `json:"labels"`
	Selector          string `json:"selector"`
	Type              string `json:"type"`
	ClusterIP         string `json:"clusterIP"`
	Ports             string `json:"ports"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type ServiceListReq struct {
	Namespace string `json:"namespace"`
}

type ServiceListResp struct {
	ListData []*ServiceListData `json:"listData"`
}
