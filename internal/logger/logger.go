package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"log"
	"os"
)

type Log struct {
	logger logrus.Logger
}

func New() *Log {
	return &Log{logger: *logrus.New()}
}
func (l *Log) Info(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *Log) Debug(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *Log) Error(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *Log) Println(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *Log) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

// Sets logs path to path from arg. If error to stdout
func (l *Log) SetPath(logPath string) {
	outFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		l.logger.SetOutput(os.Stdout)
		log.Printf("Can't find path %s. logger set to stdout", logPath)
	} else {
		l.logger.SetOutput(outFile)
	}
}

func (l *Log) SetLevel(level string) {
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalln(err)
	}
	l.logger.SetLevel(parseLevel)
}

func (l *Log) SetFormatter() {
	l.logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%",
	})
}
