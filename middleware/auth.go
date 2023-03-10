/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    auth
	@Date    2022/4/20 21:16
	@Desc
*/

package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/request"
	jwtAuth "github.com/heshaofeng1991/common/util/auth"
	"github.com/heshaofeng1991/common/util/env"
	httperr "github.com/heshaofeng1991/common/util/httpresponse"
	"github.com/sirupsen/logrus"
)

func addAuthMiddleware(router *chi.Mux) {
	router.Use(auth)
}

func auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		if !strings.Contains(req.URL.Path, "/health-check") {
			var claims jwtAuth.WMSClaims

			jwtSecret := env.JwtSecret

			if jwtSecret == "" {
				jwtSecret = "wms"
			}

			token, err := request.ParseFromRequest(
				req,
				request.AuthorizationHeaderExtractor,
				func(token *jwt.Token) (i interface{}, e error) {
					return []byte(jwtSecret), nil
				},
				request.WithClaims(&claims),
			)
			if err != nil {
				httperr.BadRequest("parse jwt token failed", err, rsp, req)

				return
			}

			logrus.Infof("token %v", token)
			logrus.Infof("error %v", err)

			if !token.Valid {
				httperr.BadRequest("invalid jwt signature", nil, rsp, req)

				return
			}
		}

		handler.ServeHTTP(rsp, req)
	})
}
