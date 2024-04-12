package handler

import (
	"net/http"

	"github.com/feihua/k8s-api/internal/logic/serve"
	"github.com/feihua/k8s-api/internal/svc"
	"github.com/feihua/k8s-api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
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
