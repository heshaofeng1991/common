package logs

import (
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Init() {
	logrus.SetFormatter(
		&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "severity",
				logrus.FieldKeyMsg:   "message",
			},
		},
	)

	if os.Getenv("LOG_FORMAT") == "text" {
		logrus.SetFormatter(
			&prefixed.TextFormatter{
				ForceFormatting: true,
			},
		)
	}

	logrus.SetLevel(logrus.DebugLevel)
}
