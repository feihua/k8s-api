package jenkins

import (
	"net/http"

	"github.com/feihua/k8s-api/internal/logic/jenkins"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
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
