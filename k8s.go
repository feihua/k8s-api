package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s_test/internal/config"
	"k8s_test/internal/handler"
	"k8s_test/internal/svc"
	"k8s_test/internal/utils/ws"
	"net/http"
	"time"
)

var configFile = flag.String("f", "etc/k8s-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.DisableStat()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			ws.WsHandler(w, r, ctx)
		},
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/test",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			ws.WsHandler(w, r, ctx)
		},
	})

	//watchlist := cache.NewListWatchFromClient(ctx.ClientSet.CoreV1().RESTClient(), "services", v1.NamespaceDefault,
	//	fields.Everything())
	//
	//_, controller := cache.NewInformer(
	//	watchlist,
	//	&v1.Service{},
	//	time.Second*0,
	//	cache.ResourceEventHandlerFuncs{
	//		AddFunc: func(obj interface{}) {
	//			fmt.Printf("service added: %s", obj)
	//		},
	//		DeleteFunc: func(obj interface{}) {
	//			fmt.Printf("service deleted: %s", obj)
	//		},
	//		UpdateFunc: func(oldObj, newObj interface{}) {
	//			fmt.Printf("service changed")
	//		},
	//	},
	//)
	//stop := make(chan struct{})
	//go controller.Run(stop)

	//watchlist := cache.NewListWatchFromClient(ctx.ClientSet.AppsV1().RESTClient(), "deployments", v1.NamespaceDefault,
	//	fields.Everything())
	//
	//_, controller := cache.NewInformer(
	//	watchlist,
	//	&v2.Deployment{},
	//	time.Second*1,
	//	cache.ResourceEventHandlerFuncs{
	//		AddFunc: func(obj interface{}) {
	//			fmt.Println("service added: %s", obj)
	//		},
	//		DeleteFunc: func(obj interface{}) {
	//			fmt.Println("service deleted: %s", obj)
	//		},
	//		UpdateFunc: func(oldObj, newObj interface{}) {
	//			fmt.Println("service changed")
	//		},
	//	},
	//)
	//stop := make(chan struct{})
	//go controller.Run(stop)

	//watchlist := cache.NewListWatchFromClient(ctx.ClientSet.ExtensionsV1beta1().RESTClient(), "ingresses", v1.NamespaceDefault,
	//	fields.Everything())
	//
	//_, controller := cache.NewInformer(
	//	watchlist,
	//	&v1beta1.Ingress{},
	//	time.Second*1,
	//	cache.ResourceEventHandlerFuncs{
	//		AddFunc: func(obj interface{}) {
	//			fmt.Println("service added: %s", obj)
	//		},
	//		DeleteFunc: func(obj interface{}) {
	//			fmt.Println("service deleted: %s", obj)
	//		},
	//		UpdateFunc: func(oldObj, newObj interface{}) {
	//			fmt.Println(oldObj)
	//			fmt.Println(newObj)
	//			fmt.Println("service changed")
	//		},
	//	},
	//)
	//stop := make(chan struct{})
	//go controller.Run(stop)

	//参考https://andblog.cn/?p=3049
	// 初始化 informer factory（为了测试方便这里设置每30s重新 List 一次）
	informerFactory := informers.NewSharedInformerFactory(ctx.ClientSet, time.Second*60)
	// 对 Deployment 监听
	deployInformer := informerFactory.Apps().V1().Deployments()
	// 创建 Informer（相当于注册到工厂中去，这样下面启动的时候就会去 List & Watch 对应的资源）
	informer := deployInformer.Informer()
	// 创建 Lister
	//deployLister := deployInformer.Lister()
	// 注册事件处理程序
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})

	// 监听 pod 资源
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAddPod,
		UpdateFunc: onUpdatePod,
		DeleteFunc: onDeletePod,
	})
	// 监听 service 资源
	servicesInformer := informerFactory.Core().V1().Services()
	servicesInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAddService,
		UpdateFunc: onUpdateService,
		DeleteFunc: onDeleteService,
	})

	// 监听 node 资源
	informerFactory.Core().V1().Nodes().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAddNode,
		UpdateFunc: onUpdateNode,
		DeleteFunc: onDeleteNode,
	})

	// 监听 namespace 资源
	informerFactory.Core().V1().Namespaces().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAddNameSpace,
		UpdateFunc: onUpdateNameSpace,
		DeleteFunc: onDeleteNameSpace,
	})

	stopper := make(chan struct{})
	defer close(stopper)

	// 启动 informer，List & Watch
	informerFactory.Start(stopper)
	// 等待所有启动的 Informer 的缓存被同步
	informerFactory.WaitForCacheSync(stopper)
	// 从本地缓存中获取 default 中的所有 deployment 列表
	//deployments, err := deployLister.Deployments("default").List(labels.Everything())
	//if err != nil {
	//	panic(err)
	//}
	//for idx, deploy := range deployments {
	//	fmt.Printf("%d -> %s\n", idx+1, deploy.Name)
	//}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	<-stopper
}
func onAdd(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	fmt.Println("add a deployment:", deploy.Name)
	fmt.Println("add a deployment:", deploy)
}
func onUpdate(old, new interface{}) {
	oldDeploy := old.(*v1.Deployment)
	newDeploy := new.(*v1.Deployment)
	fmt.Println("update deployment:", oldDeploy.Name, newDeploy.Name)
}

func onDelete(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	fmt.Println("delete a deployment:", deploy.Name)
}
func onAddPod(obj interface{}) {
	deploy := obj.(*v2.Pod)
	fmt.Println("add a pod:", deploy.Name)
}
func onUpdatePod(old, new interface{}) {
	oldDeploy := old.(*v2.Pod)
	newDeploy := new.(*v2.Pod)
	fmt.Println("update pod:", oldDeploy.Name, newDeploy.Name)
}

func onDeletePod(obj interface{}) {
	deploy := obj.(*v2.Pod)
	fmt.Println("delete a pod:", deploy.Name)
}
func onAddService(obj interface{}) {
	deploy := obj.(*v2.Service)
	fmt.Println("add a service:", deploy.Name)
}
func onUpdateService(old, new interface{}) {
	oldDeploy := old.(*v2.Service)
	newDeploy := new.(*v2.Service)
	fmt.Println("update service:", oldDeploy.Name, newDeploy.Name)
}

func onDeleteService(obj interface{}) {
	deploy := obj.(*v2.Service)
	fmt.Println("delete a service:", deploy.Name)
}
func onAddNode(obj interface{}) {
	deploy := obj.(*v2.Node)
	fmt.Println("add a node:", deploy.Name)
}
func onUpdateNode(old, new interface{}) {
	oldDeploy := old.(*v2.Node)
	newDeploy := new.(*v2.Node)
	fmt.Println("update node:", oldDeploy.Name, newDeploy.Name)
}

func onDeleteNode(obj interface{}) {
	deploy := obj.(*v2.Node)
	fmt.Println("delete a node:", deploy.Name)
}
func onAddNameSpace(obj interface{}) {
	deploy := obj.(*v2.Namespace)
	fmt.Println("add a namespace:", deploy.Name)
}
func onUpdateNameSpace(old, new interface{}) {
	oldDeploy := old.(*v2.Namespace)
	newDeploy := new.(*v2.Namespace)
	fmt.Println("update namespace:", oldDeploy.Name, newDeploy.Name)
}

func onDeleteNameSpace(obj interface{}) {
	deploy := obj.(*v2.Namespace)
	fmt.Println("delete a namespace:", deploy.Name)
}
