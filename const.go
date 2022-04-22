/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    const
	@Date    2022/4/13 23:42
	@Desc
*/

package internal

// Sentry DSN.

const (
	SentrySyncTimeout = 3
)

// HTTP.

const (
	CorsMaxAge      = 300
	EntryTimeLength = 100
	SentryTimeout   = 5
)

const (
	StackSize = 2048
)

type UserKey string

const (
	UserID UserKey = "UserID"
)

const (
	PlatformJingDong = "JingDong"
	Format           = "2006-01-02 15:04:05"
	NSS              = "NSS"
)

const (
	JDInboundStatus          = 30
	NSSInboundStatus         = 30
	NSSOutboundOrderShipped  = 50
	JDOutboundOrderShipped   = 40
	NSSOutboundOrderFinished = 80
)

const (
	Success = "success"
)
