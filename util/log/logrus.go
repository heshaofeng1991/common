/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    logrus
	@Date    2022/4/14 15:43
	@Desc
*/

package log

import (
	"fmt"
	"net/http"
	"os"
	"time"

	internal "github.com/NextSmartShip/common"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// InitLog 初始化 loggers.
func InitLog() {
	// 设置日志格式为json格式.
	log.SetFormatter(
		&log.JSONFormatter{
			TimestampFormat:  "2006-01-02 15:04:05",
			PrettyPrint:      true,
			DisableTimestamp: false,
			DataKey:          "",
			FieldMap:         nil,
		},
	)

	if os.Getenv("LOG_FORMAT") == "text" {
		log.SetFormatter(
			&prefixed.TextFormatter{
				ForceFormatting: true,
			},
		)
	}

	// 设置只记录日志级别为trace及其以上的日志.
	log.SetLevel(log.TraceLevel)

	// 设置日志输出等级 stderr stdout.
	log.SetOutput(os.Stdout)

	// 设置输出文件名，行号和函数名.
	log.SetReportCaller(true)
}

func NewStructuredLogger(logger *log.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{logger})
}

type StructuredLogger struct {
	Logger *log.Logger
}

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{Logger: log.NewEntry(l.Logger)}
	logFields := log.Fields{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		logFields["req_id"] = reqID
	}

	logFields["http_method"] = r.Method
	logFields["remote_addr"] = r.RemoteAddr
	logFields["uri"] = r.RequestURI
	logFields["path"] = r.URL.Path

	entry.Logger = entry.Logger.WithFields(logFields)

	entry.Logger.Info("Request started")

	return entry
}

type StructuredLoggerEntry struct {
	Logger log.FieldLogger
}

func (l *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Logger = l.Logger.WithFields(
		log.Fields{
			"resp_status": status, "resp_bytes_length": bytes,
			"resp_elapsed": elapsed.Round(time.Millisecond / internal.EntryTimeLength).String(),
		},
	)

	l.Logger.Info("Request completed	")
}

func (l *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger = l.Logger.WithFields(
		log.Fields{
			"stack": string(stack),
			"panic": fmt.Sprintf("%+v", v),
		},
	)
}

func GetLogEntry(r *http.Request) log.FieldLogger {
	entry, ok := middleware.GetLogEntry(r).(*StructuredLoggerEntry)
	if !ok {
		return nil
	}

	return entry.Logger
}
