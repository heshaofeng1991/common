/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    recover
	@Date    2022/4/20 20:52
	@Desc
*/

package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi/v5"
	internal "github.com/heshaofeng1991/common"
	"github.com/sirupsen/logrus"
)

func addCatchExceptionMiddleware(router *chi.Mux) {
	router.Use(catch)
}

func catch(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				if env.SentryDsn != "" {
					log.LogUtil.Errorf("openapi: %v", r)
					log.LogUtil.PrintPanicStack()
					sentry.RecoverWithContext(context.Background())
					sentry.CurrentHub().Recover(r)
					sentry.Flush(time.Second * internal.SentryTimeout)
				}

				logrus.Panicf("panic from sentry : %v", r)
			}
		}()

		handler.ServeHTTP(rsp, req.WithContext(req.Context()))
	})
}
