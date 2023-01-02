/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    middleware
	@Date    2022/4/19 14:19
	@Desc
*/

package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/heshaofeng1991/common/util/log"
	"github.com/sirupsen/logrus"
)

func SetMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(log.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)
	addCatchExceptionMiddleware(router)

	router.Use(middleware.NoCache)
}
