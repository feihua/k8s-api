package handler

import (
	"net/http"
	"k8s_test/internal/logic/pods"
	"github.com/tal-tech/go-zero/rest/httpx"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"
)

func PodLogHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PodLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPodLogLogic(r.Context(), ctx)
		resp, err := l.PodLog(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
