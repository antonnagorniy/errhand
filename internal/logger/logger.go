package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"log"
	"os"
	"sync"
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

func Println(args ...interface{}) {
	logger.Println(args...)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
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

func SetFormatter() {
	logger.SetFormatter(&prefixed.TextFormatter{
		ForceColors:      true,
		DisableColors:    false,
		ForceFormatting:  true,
		DisableTimestamp: false,
		DisableUppercase: false,
		FullTimestamp:    true,
		TimestampFormat:  "2020-01-02 15:04:05",
		DisableSorting:   false,
		QuoteEmptyFields: false,
		QuoteCharacter:   "",
		SpacePadding:     0,
		Once:             sync.Once{},
	})

}
