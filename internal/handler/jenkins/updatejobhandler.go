package jenkins

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8s_test/internal/logic/jenkins"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

func UpdateJobHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateJenkinsJobReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := jenkins.NewUpdateJobLogic(r.Context(), ctx)
		resp, err := l.UpdateJob(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
