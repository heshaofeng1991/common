/*
	@Author  johnny
	@Author  heshaofeng1991@gmail.com
	@Version v1.0.0
	@File    logrus
	@Date    2022/4/13 23:41
	@Desc
*/

package sentry

import (
	"time"

	internal "github.com/heshaofeng1991/common"
	"github.com/heshaofeng1991/common/util/env"
)

// Init 初始化sentry.
func Init() {
	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	sentrySyncTransport.Timeout = time.Second * internal.SentrySyncTimeout

	_ = sentry.Init(sentry.ClientOptions{
		Dsn:       env.SentryDsn,
		Transport: sentrySyncTransport,
	})
}
