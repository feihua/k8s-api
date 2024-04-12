package jenkins

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8s_test/internal/logic/jenkins"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

func AddJobToViewHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddJobToViewReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := jenkins.NewAddJobToViewLogic(r.Context(), ctx)
		resp, err := l.AddJobToView(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
