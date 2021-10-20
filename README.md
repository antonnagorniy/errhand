# Simple errors handler and logger
## Based on logrus. 
#### It's just a simple decorator for quick logger config and errors handling.

[![Actions Status](https://github.com/kattaris/errhand/workflows/master CI/badge.svg)](https://github.com/kattaris/errhand/actions)
[![Coverage Status](https://codecov.io/github/kattaris/errhand/coverage.svg?branch=master)](https://codecov.io/gh/kattaris/errhand)
[![Releases](https://img.shields.io/github/v/release/errhand/errhand.svg?include_prereleases&style=flat-square)](https://github.com/kattaris/errhand/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

### Custom log and handler parameters:
```
func init() {
	userName := os.Getenv("USER")
	logs := fmt.Sprintf("/home/%s/Logs/errhand/main.log", userName)
	errhand.Init(logs)
}

func main() {
	log := errhand.GetLogger()
	err := errors.New("test error")

	log.HandleError(err, false)
	log.Infoln("Custom info")
	log.Debugln("Custom debug")
	log.Warnln("Custom warn")
	log.Fatalln("Custom fatal")
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
