# Simple errors handler and logger
## Based on logrus. 
#### It's just a simple decorator for quick logger config and errors handling.

[![Actions Status](https://github.com/kattaris/errhand/workflows/CI/badge.svg)](https://github.com/kattaris/errhand/actions)
[![Coverage Status](https://codecov.io/github/kattaris/errhand/coverage.svg?branch=master)](https://codecov.io/gh/kattaris/errhand)
[![Releases](https://img.shields.io/github/v/release/errhand/errhand.svg?include_prereleases&style=flat-square)](https://github.com/kattaris/errhand/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

### Custom log and handler parameters:
```
var log *errhand.Errhand

func init() {
	log = errhand.GetLogger()
}

func main() {
	err := errors.New("test")
	logger.HandleError(err, false)
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
