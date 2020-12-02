package errhand

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"sync"
)

type Errhand struct {
	Log *logrus.Logger
}

func New() *Errhand {
	return &Errhand{Log: logrus.New()}
}

// HandleError handles simple errors
func (e *Errhand) HandleError(err error) {
	if err == nil {
		return
	}

	cause := errors.Cause(err)
	switch cause.(type) {
	case error:
		e.Log.Errorln(err)
	default:
		e.Log.Panicln(err)
	}

}

func (e *Errhand) Infoln(v ...interface{}) {
	e.Log.Infoln(v...)
}

func (e *Errhand) Infof(format string, v ...interface{}) {
	e.Log.Infof(format, v...)
}

func (e *Errhand) Errorln(v ...interface{}) {
	e.Log.Errorln(v...)
}

func (e *Errhand) Errorf(format string, v ...interface{}) {
	e.Log.Errorf(format, v...)
}

func (e *Errhand) Debugln(v ...interface{}) {
	e.Log.Debugln(v...)
}

func (e *Errhand) Debugf(format string, v ...interface{}) {
	e.Log.Debugf(format, v...)
}

func (e *Errhand) Warnln(v ...interface{}) {
	e.Log.Warnln(v...)
}

func (e *Errhand) Warnf(format string, v ...interface{}) {
	e.Log.Warnf(format, v...)
}

func (e *Errhand) Fatalln(v ...interface{}) {
	e.Log.Fatalln(v...)
}

func (e *Errhand) Fatalf(format string, v ...interface{}) {
	e.Log.Fatalf(format, v...)
}

func (e *Errhand) WithError(err error) *logrus.Entry {
	return e.Log.WithError(err)
}

func (e *Errhand) WithField(key string, value interface{}) *logrus.Entry {
	return e.Log.WithField(key, value)
}

// Set custom output and log level
func (e *Errhand) CustomLogger(logsPath string, level string) {
	e.setPath(logsPath)
	e.setLevel(level)
	e.Log.SetFormatter(&prefixed.TextFormatter{
		ForceColors:      false,
		DisableColors:    true,
		ForceFormatting:  true,
		DisableTimestamp: false,
		DisableUppercase: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableSorting:   false,
		QuoteEmptyFields: false,
		QuoteCharacter:   "",
		SpacePadding:     0,
		Once:             sync.Once{},
	})
}

// Set custom output
func (e *Errhand) setPath(logPath string) {
	outFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		e.Log.SetOutput(os.Stdout)
		e.Log.Printf("Can't find path %s. Log set to stdout", logPath)
	} else {
		e.Log.SetOutput(outFile)
	}
}

// Set custom log level
func (e *Errhand) setLevel(level string) {
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		e.Log.Fatalln(err)
	}
	e.Log.SetLevel(parseLevel)
}
