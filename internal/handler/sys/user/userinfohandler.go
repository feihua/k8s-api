package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"k8s_test/internal/logic/sys/user"
	"k8s_test/internal/svc"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
