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
	"bufio"
	"errors"
	"os"
	"strings"
)

var envMap map[string]string

func GetEnvFromFile(key string) (string, bool) {
	if len(envMap) == 0 {
		if err := ReadEnvFile(".env"); err != nil {
			return "", false
		}
	}

	if value, ok := envMap[key]; ok && value != "" {
		return value, true
	}

	return "", false
}

func ReadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	envMap = make(map[string]string)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return err
	}
	for _, fullLine := range lines {
		if !isIgnoredLine(fullLine) {
			var key, value string
			key, value, err = parseLine(fullLine)
			if err != nil {
				return err
			}
			envMap[key] = value
		}
	}
	return nil
}

func isIgnoredLine(line string) bool {
	trimmedLine := strings.TrimSpace(line)
	return len(trimmedLine) == 0 || strings.HasPrefix(trimmedLine, "#")
}

func parseLine(line string) (key string, value string, err error) {
	if len(line) == 0 {
		err = errors.New("zero length string")
		return
	}

	splitString := strings.SplitN(line, "=", 2)
	if len(splitString) != 2 {
		err = errors.New("Can't separate key from value")
		return
	}

	key = strings.TrimSpace(splitString[0])
	value = strings.Trim(splitString[1], " ")
	return
}

func getEnv(key string, defaultValue string) string {
	if value, ok := GetEnvFromFile(key); ok {
		return value
	}

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
