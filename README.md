# Simple errors handler and logger
## Based on logrus. It's just a simple decorator for quick logger config and errors handling.

### Using erorrs handler:
```
var logger = errhand.Errhand{}

err := errors.New("test")
logger.HandleSimpleErr(err)
```

### Custom log and handler parameters:
```
var logger = errhand.Errhand{}
var logsPath = "~/logs/testLogs.log"

func init() {
	logger.CustomLogger(logsPath, "debug")
}

func main() {
	err := errors.New("test")
	logger.HandleSimpleErr(err)
	logger.Infoln("Custom info")
	logger.Debugln("Custom debug")
}
```

For logs and handler output to stdout don't use CustomLogger method.
