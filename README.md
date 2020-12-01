# Simple errors handler and logger
## Based on logrus. It's just a simple decorator for quick logger config and errors handling.

### Using erorrs handler:
```
var log = errhand.New()

func main() {
    err := errors.New("test")
    log.HandleSimpleErr(err)
}
```
You'll get smth like that in stdOut:
```
time="2020-12-01T15:35:27+03:00" level=error msg=test
```

### Custom log and handler parameters:
```
var log = errhand.New()

func init() {
	logsPath := "~/logs/testLogs.log"
	logger.CustomLogger(logsPath, "debug")
}

func main() {
	err := errors.New("test")
	logger.HandleSimpleErr(err)
	logger.Infoln("Custom info")
	logger.Debugln("Custom debug")
}
```
You'll get smth like that in specified log file:
```
[2020-12-01 15:47:43] ERROR test
[2020-12-01 15:47:43]  INFO Custom info
[2020-12-01 15:47:43] DEBUG Custom debug
```
