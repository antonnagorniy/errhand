package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var logger = logrus.New()

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func SetPath(sysVarLogPath string) {
	if value, exists := os.LookupEnv(sysVarLogPath); exists {
		outFile, err := os.OpenFile(value, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalln(err)
		}
		logger.SetOutput(outFile)
	} else {
		logger.SetOutput(os.Stdout)
	}
}

func SetLevel(level string) {
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalln(err)
	}
	logger.SetLevel(parseLevel)
}
