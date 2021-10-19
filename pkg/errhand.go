package errhand

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
		if err != nil {
			return err
		}
	}

	return err
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

// Errhand main struct
type Errhand struct {
	*logrus.Entry
}

// GetLogger returns fully initialized logger
func GetLogger() *Errhand {
	return &Errhand{e}
}

// Init initializes logger with custom logs path
func Init(logdir string) {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &prefixed.TextFormatter{
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
	}

	file, err := createLogFile(logdir)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writeHook{
		Writer:    []io.Writer{file, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(l)
}

// HandleError handles simple errors
func (e *Errhand) HandleError(err error, fatal bool) {
	if err == nil {
		return
	}

	if fatal {
		e.Fatalln(err)
	}

	e.Errorln(err)
}

// Create directories if not exists
func createLogFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0770); err != nil {
		return nil, err
	} else if _, errFile := os.Stat(path); os.IsNotExist(errFile) {
		file, err2 := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err2 != nil {
			return nil, err2
		}
		return file, nil
	} else {
		file, err3 := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
		if err3 != nil {
			return nil, err3
		}
		return file, nil
	}
}
