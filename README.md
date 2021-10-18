# Simple errors handler and logger
## Based on logrus. 
#### It's just a simple decorator for quick logger config and errors handling.

### Custom log and handler parameters:
```
var log *errhand.Errhand

func init() {
	log = errhand.GetLogger()
}

func main() {
	err := errors.New("test")
	logger.HandleSimpleErr(err)
	logger.Infoln("Custom info")
	logger.Debugln("Custom debug")
	logger.Warnln("Custom warn")
	logger.Fatalln("Custom fatal")
}
```
You'll get smth like that in specified log file:
```
[2020-12-01 16:03:06] ERROR test
[2020-12-01 16:03:06]  INFO Custom info
[2020-12-01 16:03:06] DEBUG Custom debug
[2020-12-01 16:03:06]  WARN Custom warn
[2020-12-01 16:03:06] FATAL Custom warn
```
