package svc

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	ClientSet *kubernetes.Clientset
	DbEngine  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	kubeConfig := "etc/config"
	buildConfigFromFlags, err := clientcmd.BuildConfigFromFlags("", kubeConfig)

	if err != nil {
		logx.Errorf("获取kubeConfig：%s异常: %s", kubeConfig, err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(buildConfigFromFlags)
	if err != nil {
		logx.Errorf("连接kubernetes：%s异常: %s", kubeConfig, err.Error())
	}

	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tech_",
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
	//自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	//db.AutoMigrate(&models.User{})

	return &ServiceContext{
		Config:    c,
		ClientSet: clientSet,
		DbEngine:  db,
	}
}

func (sc ServiceContext) NodeList() (*v1.NodeList, error) {
	return sc.ClientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
}
func (sc ServiceContext) NodeGet(name string) (*v1.Node, error) {
	return sc.ClientSet.CoreV1().Nodes().Get(context.TODO(), name, metaV1.GetOptions{})
}
