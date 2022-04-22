/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    logrus
	@Date    2022/4/13 23:41
	@Desc
*/

package sentry

import (
	"time"

	internal "github.com/NextSmartShip/common"
	"github.com/NextSmartShip/common/util/env"
	"github.com/getsentry/sentry-go"
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
