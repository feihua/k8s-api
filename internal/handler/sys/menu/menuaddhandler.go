package handler

import (
	"net/http"

	"k8s_test/internal/logic/sys/menu"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MenuAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMenuAddLogic(r.Context(), ctx)
		resp, err := l.MenuAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
