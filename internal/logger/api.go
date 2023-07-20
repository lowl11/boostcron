package logger

import (
	"github.com/lowl11/lazylog/log"
	"github.com/lowl11/lazylog/log/log_internal"
	systemLog "log"
)

func Error(err error, message string) {
	if log_internal.Initialized() {
		log.Error(err, message)
	} else {
		var logMessage string
		if err != nil {
			logMessage += err.Error() + " -> "
		}

		logMessage += message
		systemLog.Println(logMessage)
	}
}
