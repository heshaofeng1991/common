package log

import (
	"runtime"

	internal "github.com/NextSmartShip/wms-backend/common"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

var LogUtil = &logUtil{}

type logUtil struct{}

func (l *logUtil) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)

	sentry.CaptureException(internal.Stack(format, args))
}

func (l *logUtil) PrintPanicStack() {
	// 打印调用栈信息.
	buffer := make([]byte, internal.StackSize)

	bufferSize := runtime.Stack(buffer, false)

	stackInfo := string(buffer[:bufferSize])

	logrus.Errorf("panic stack info: %s", stackInfo)
}
