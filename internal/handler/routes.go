// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	configmap "github.com/feihua/k8s-api/internal/handler/configmap"
	deployment "github.com/feihua/k8s-api/internal/handler/deployment"
	ingress "github.com/feihua/k8s-api/internal/handler/ingress"
	jenkins "github.com/feihua/k8s-api/internal/handler/jenkins"
	namespace "github.com/feihua/k8s-api/internal/handler/namespace"
	nodes "github.com/feihua/k8s-api/internal/handler/nodes"
	pods "github.com/feihua/k8s-api/internal/handler/pods"
	secret "github.com/feihua/k8s-api/internal/handler/secret"
	serve "github.com/feihua/k8s-api/internal/handler/serve"
	statefulset "github.com/feihua/k8s-api/internal/handler/statefulset"
	sysdept "github.com/feihua/k8s-api/internal/handler/sys/dept"
	sysdict "github.com/feihua/k8s-api/internal/handler/sys/dict"
	sysjob "github.com/feihua/k8s-api/internal/handler/sys/job"
	syslog "github.com/feihua/k8s-api/internal/handler/sys/log"
	sysmenu "github.com/feihua/k8s-api/internal/handler/sys/menu"
	sysrole "github.com/feihua/k8s-api/internal/handler/sys/role"
	sysuser "github.com/feihua/k8s-api/internal/handler/sys/user"
	"github.com/feihua/k8s-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/create",
				Handler: deployment.DeploymentCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/update",
				Handler: deployment.DeploymentUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/updateStatus",
				Handler: deployment.DeploymentUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/delete",
				Handler: deployment.DeploymentDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/deleteCollection",
				Handler: deployment.DeploymentDeleteCollectionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/get",
				Handler: deployment.DeploymentGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/list",
				Handler: deployment.DeploymentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/getScale",
				Handler: deployment.DeploymentGetScaleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/deployment/updateScale",
				Handler: deployment.DeploymentUpdateScaleHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/ingress/create",
				Handler: ingress.IngressCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/ingress/update",
				Handler: ingress.IngressUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/ingress/updateStatus",
				Handler: ingress.IngressUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/ingress/delete",
				Handler: ingress.IngressDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/ingress/get",
				Handler: ingress.IngressGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/ingress/list",
				Handler: ingress.IngressListHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/namespace/create",
				Handler: namespace.NamespaceCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/namespace/delete",
				Handler: namespace.NamespaceDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/namespace/get",
				Handler: namespace.NamespaceGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/namespace/list",
				Handler: namespace.NamespaceListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/namespace/listWith",
				Handler: namespace.NamespaceListWithHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/service/create",
				Handler: serve.ServiceCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/update",
				Handler: serve.ServiceUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/updateStatus",
				Handler: serve.ServiceUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/delete",
				Handler: serve.ServiceDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/get",
				Handler: serve.ServiceGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/list",
				Handler: serve.ServiceListHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/create",
				Handler: pods.PodCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/update",
				Handler: pods.PodUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/updateStatus",
				Handler: pods.PodUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/delete",
				Handler: pods.PodDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/deleteCollection",
				Handler: pods.PodDeleteCollectionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/get",
				Handler: pods.PodGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/list",
				Handler: pods.PodListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/getEphemeralContainers",
				Handler: pods.PodGetEphemeralContainersHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/updateEphemeralContainers",
				Handler: pods.PodUpdateEphemeralContainersHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/pods/log",
				Handler: pods.PodLogHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/nodes/get",
				Handler: nodes.NodeGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/nodes/list",
				Handler: nodes.NodeListHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/statefulset/create",
				Handler: statefulset.StatefulSetCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/statefulset/update",
				Handler: statefulset.StatefulSetUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/statefulset/updateStatus",
				Handler: statefulset.StatefulSetUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/statefulset/delete",
				Handler: statefulset.StatefulSetDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/statefulset/get",
				Handler: statefulset.StatefulSetGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/statefulset/list",
				Handler: statefulset.StatefulSetListHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/secret/create",
				Handler: secret.SecretCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/secret/update",
				Handler: secret.SecretUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/secret/updateStatus",
				Handler: secret.SecretUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/secret/delete",
				Handler: secret.SecretDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/secret/get",
				Handler: secret.SecretGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/secret/list",
				Handler: secret.SecretListHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/configMap/create",
				Handler: configmap.ConfigMapCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/configMap/update",
				Handler: configmap.ConfigMapUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/configMap/updateStatus",
				Handler: configmap.ConfigMapUpdateStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/configMap/delete",
				Handler: configmap.ConfigMapDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/configMap/get",
				Handler: configmap.ConfigMapGetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/configMap/list",
				Handler: configmap.ConfigMapListHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/sys/user/currentUser",
				Handler: sysuser.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/add",
				Handler: sysuser.UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/list",
				Handler: sysuser.UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/update",
				Handler: sysuser.UserUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/delete",
				Handler: sysuser.UserDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/reSetPassword",
				Handler: sysuser.ReSetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/UpdateUserStatus",
				Handler: sysuser.UpdateUserStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/selectAllData",
				Handler: sysuser.SelectAllDataHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/login",
				Handler: sysuser.UserLoginHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/role/add",
				Handler: sysrole.RoleAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/role/list",
				Handler: sysrole.RoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/role/update",
				Handler: sysrole.RoleUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/role/delete",
				Handler: sysrole.RoleDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/role/roleMenuIds",
				Handler: sysrole.RoleMenuIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/role/queryMenuByRoleId",
				Handler: sysrole.QueryMenuByRoleIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/role/updateRoleMenu",
				Handler: sysrole.UpdateRoleMenuHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/menu/add",
				Handler: sysmenu.MenuAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/menu/list",
				Handler: sysmenu.MenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/menu/update",
				Handler: sysmenu.MenuUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/menu/delete",
				Handler: sysmenu.MenuDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dict/add",
				Handler: sysdict.DictAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dict/list",
				Handler: sysdict.DictListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dict/update",
				Handler: sysdict.DictUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dict/delete",
				Handler: sysdict.DictDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dept/add",
				Handler: sysdept.DeptAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dept/list",
				Handler: sysdept.DeptListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dept/update",
				Handler: sysdept.DeptUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/dept/delete",
				Handler: sysdept.DeptDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/job/add",
				Handler: sysjob.JobAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/job/list",
				Handler: sysjob.JobListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/job/update",
				Handler: sysjob.JobUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/job/delete",
				Handler: sysjob.JobDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/loginLog/list",
				Handler: syslog.LoginLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/loginLog/delete",
				Handler: syslog.LoginLogDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/list",
				Handler: jenkins.GetAllJobsNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/info",
				Handler: jenkins.GetInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/createJob",
				Handler: jenkins.CreateJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/updateJob",
				Handler: jenkins.UpdateJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/renameJob",
				Handler: jenkins.RenameJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/copyJob",
				Handler: jenkins.CopyJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/deleteJob",
				Handler: jenkins.DeleteJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/buildJob",
				Handler: jenkins.BuildJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/getBuild",
				Handler: jenkins.GetBuildHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/getJob",
				Handler: jenkins.GetJobHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/getAllJobs",
				Handler: jenkins.GetAllJobsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/getView",
				Handler: jenkins.GetViewHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/getAllViews",
				Handler: jenkins.GetAllViewsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/getAllViewsWith",
				Handler: jenkins.GetAllViewsWithHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/createView",
				Handler: jenkins.CreateViewHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/addJobToView",
				Handler: jenkins.AddJobToViewHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/jenkins/getAllBuildIds",
				Handler: jenkins.GetAllBuildIdsHandler(serverCtx),
			},
		},
	)
}
