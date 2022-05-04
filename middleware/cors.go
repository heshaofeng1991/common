/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    cors
	@Date    2022/4/19 14:20
	@Desc
*/

package middleware

import (
	internal "github.com/NextSmartShip/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

func addCorsMiddleware(router *chi.Mux) {
	corsMiddleware := cors.New(cors.Options{
		AllowedMethods: []string{
			"GET",
			"POST",
			"OPTIONS",
			"PUT",
			"DELETE",
			"PATCH",
		},
		AllowedHeaders: []string{
			"X-Requested-With",
			"Content-Type",
			"Accept",
			"Origin",
			"Authorization",
			"X-Api-Version",
			"x-nss-tenant-id",
			"Access-Control-Allow-Origin",
		},
		AllowCredentials: true,
		MaxAge:           internal.CorsMaxAge,
		AllowOriginFunc: AllowOriginFunc,
	})

	router.Use(corsMiddleware.Handler)
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	origin = r.Header.Get("Origin")
	
	if origin == "*" {
		return true
	}
	
	return false
}