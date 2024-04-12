package jenkins

import (
	"net/http"

	"github.com/feihua/k8s-api/internal/logic/jenkins"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := jenkins.NewGetInfoLogic(r.Context(), ctx)
		resp, err := l.GetInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
