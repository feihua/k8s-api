package handler

import (
	"net/http"

	"github.com/feihua/k8s-api/internal/logic/sys/job"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func JobUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateJobReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewJobUpdateLogic(r.Context(), ctx)
		resp, err := l.JobUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
