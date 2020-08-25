package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"log"
	"os"
)

var logger = logrus.New()

func Info(args ...interface{}) {
	logger.Infoln(args...)
}

func Debug(args ...interface{}) {
	logger.Debugln(args...)
}

func Error(args ...interface{}) {
	logger.Errorln(args...)
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
	logger.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%",
	})

}
