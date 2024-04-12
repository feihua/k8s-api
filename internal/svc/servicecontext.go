package svc

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s_test/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	ClientSet *kubernetes.Clientset
	DbClient  *gorm.DB
	RestConf  *rest.Config
	Jenkins   *gojenkins.Jenkins
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

	logx.Info("连接kubernetes成功")

	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.Mysql.Datasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	logx.Info("连接Mysql成功")
	//自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	//db.AutoMigrate(&models.User{})

	jenkins := gojenkins.CreateJenkins(nil, "http://10.168.12.132:9093/jenkins", "admin", "admin@UAT18")

	ctx := context.Background()
	_, err = jenkins.Init(ctx)
	if err != nil {
		logx.Errorf("连接Jenkins失败: %s", err.Error())
	}

	logx.Info("连接Jenkins成功")

	return &ServiceContext{
		Config:    c,
		ClientSet: clientSet,
		DbClient:  db,
		RestConf:  buildConfigFromFlags,
		Jenkins:   jenkins,
	}
}

func (sc ServiceContext) NodeList() (*v1.NodeList, error) {
	return sc.ClientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
}
func (sc ServiceContext) NodeGet(name string) (*v1.Node, error) {
	return sc.ClientSet.CoreV1().Nodes().Get(context.TODO(), name, metaV1.GetOptions{})
}
