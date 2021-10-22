package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"k8s_test/internal/config"
	"k8s_test/internal/handler"
	"k8s_test/internal/svc"
	"k8s_test/internal/utils/ws"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
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

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
