package middleware

import (
	"net/http"

	"github.com/Zeroaril7/go-rest-study/helper"
	"github.com/Zeroaril7/go-rest-study/model/web/response"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "palasdweij" == r.Header.Get("X-API-KEY") {
		m.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		res := response.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResBody(w, res)
	}
}
