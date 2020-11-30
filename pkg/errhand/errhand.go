package errhand

import (
	"github.com/kattaris/errhand/internal/logger"
)

var customLog = logger.New()

type Handler struct {
	log *logger.Log
}

// HandleSimpleErr handles simple errors
func (h *Handler) HandleSimpleErr(err error) {
	if err != nil {
		h.log.Error(err, "\n")
	}
}

// Return Handler with log path and level.
func New(logPath string, level string) *Handler {
	handler := Handler{log: customLog}
	handler.log.SetPath(logPath)
	handler.log.SetLevel(level)
	handler.log.SetFormatter()
	return &handler
}

// Print with new line
func (h *Handler) Println(v ...interface{}) {
	h.log.Println(v...)
}

// Print with format
func (h *Handler) Printf(format string, v ...interface{}) {
	h.log.Printf(format, v...)
}
