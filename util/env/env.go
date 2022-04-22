/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    env
	@Date    2022/4/13 23:41
	@Desc
*/

package env

import (
	"os"
)

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

var SentryDsn = getEnv("SENTRY_DSN", "https://0e9afcc145234e5aa2ab9a7780fe62af@o504873.ingest.sentry.io/6141414")

var ServerType = getEnv("SERVER_TO_RUN", "http")

var ServerPort = getEnv("SERVER_PORT", "3001")

var MysqlDSN = getEnv("MYSQL_DSN", "root:@tcp(127.0.0.1:3306)/test")

var JwtSecret = getEnv("JWT_SECRET", "wms")

var AwsRegion = getEnv("AWS_REGION", "ap-east-1")

var AwsAccessKeyID = getEnv("AWS_ACCESS_KEY_ID", "AKIA4YTYMDHLOT4P7NM5")

var AwsSecretAccessKey = getEnv("AWS_SECRET_ACCESS_KEY", "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")

var AwsWmsBackendQueueURL = getEnv("AWS_WMS_BACKEND_QUEUE_URL",
	"https://sqs.ap-east-1.amazonaws.com/877499521494/dev-wms-backend")
