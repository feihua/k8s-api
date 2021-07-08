package handler

import (
	"net/http"

	"k8s_test/internal/logic/statefulset"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func StatefulSetWatchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatefulSetWatchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStatefulSetWatchLogic(r.Context(), ctx)
		resp, err := l.StatefulSetWatch(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
