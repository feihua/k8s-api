package handler

import (
	"net/http"

	"k8s_test/internal/logic/serve"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ServiceUpdateStatusHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ServiceUpdateStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewServiceUpdateStatusLogic(r.Context(), ctx)
		resp, err := l.ServiceUpdateStatus(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
