package handler

import (
	"github.com/kattaris/errhand/v2/internal/logger"
)

type Handler struct {
}

// HandleSimpleErr handles simple errors
func (Handler) HandleSimpleErr(err error, message string) {
	if err != nil {
		logger.Error(message, err)
	}
}

func New(sysVarLogPath string, level string) Handler {
	logger.SetPath(sysVarLogPath)
	logger.SetLevel(level)
	return Handler{}
}

func Println(v ...interface{}) {
	logger.Println(v...)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v)
}
