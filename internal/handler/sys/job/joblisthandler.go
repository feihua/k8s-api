package handler

import (
	"net/http"

	"k8s_test/internal/logic/sys/job"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func JobListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListJobReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewJobListLogic(r.Context(), ctx)
		resp, err := l.JobList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
