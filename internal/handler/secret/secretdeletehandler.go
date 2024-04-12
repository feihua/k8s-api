package handler

import (
	"net/http"

	"k8s_test/internal/logic/secret"
	"k8s_test/internal/svc"
	"k8s_test/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SecretDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SecretDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSecretDeleteLogic(r.Context(), ctx)
		resp, err := l.SecretDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
