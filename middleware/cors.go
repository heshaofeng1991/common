/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    cors
	@Date    2022/4/19 14:20
	@Desc
*/

package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	internal "github.com/heshaofeng1991/common"
	"github.com/rs/cors"
)

func addCorsMiddleware(router *chi.Mux) {
	corsMiddleware := cors.New(
		cors.Options{
			AllowedMethods:         []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
			AllowedHeaders:         []string{"authorization", "content-type", "x-csrf-token", "x-requested-with"},
			AllowCredentials:       true,
			MaxAge:                 internal.CorsMaxAge,
			AllowOriginRequestFunc: AllowOriginRequestFunc,
			OptionsSuccessStatus:   http.StatusNoContent,
		},
	)

	router.Use(corsMiddleware.Handler)
}

func AllowOriginRequestFunc(r *http.Request, origin string) bool {
	return origin == r.Header.Get("Origin")
}
