package errhand

import (
	"github.com/kattaris/errhand/internal/logger"
)

type Handler struct {
}

// HandleSimpleErr handles simple errors
func (Handler) HandleSimpleErr(err error) {
	if err != nil {
		logger.Error(err)
	}
}

// Return Handler with log path and level
func New(sysVarLogPath string, level string) Handler {
	logger.SetPath(sysVarLogPath)
	logger.SetLevel(level)
	logger.SetFormatter()
	return Handler{}
}

// Print with new line
func (Handler) Println(v ...interface{}) {
	logger.Println(v...)
}

// Print with format
func (Handler) Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
