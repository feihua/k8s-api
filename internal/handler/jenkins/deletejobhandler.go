package jenkins

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8s_test/internal/logic/jenkins"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

func DeleteJobHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteJenkinsJobReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := jenkins.NewDeleteJobLogic(r.Context(), ctx)
		resp, err := l.DeleteJob(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
