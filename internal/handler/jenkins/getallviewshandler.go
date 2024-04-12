package jenkins

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8s_test/internal/logic/jenkins"
	"k8s_test/internal/svc"
)

func GetAllViewsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := jenkins.NewGetAllViewsLogic(r.Context(), ctx)
		resp, err := l.GetAllViews()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
