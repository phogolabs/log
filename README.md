# LOG

[![Documentation][godoc-img]][godoc-url]
![License][license-img]
[![Build Status][travis-img]][travis-url]
[![Coverage][codecov-img]][codecov-url]
[![Go Report Card][report-img]][report-url]

*A simple logger for Golang*

## Installation

```console
$ go get -u github.com/phogolabs/log
```

## Getting started

```golang
import (
  "github.com/phogolabs/log"
  "github.com/phogolabs/log/console"
)

log.SetHandler(console.New(os.Stdout))

logger := log.WithFields(log.F("app", "service-api"))
logger.Info("Hello")
```

## Levels

The levels follow similar to syslog concept:

**DebugLevel** - Info useful to developers for debugging the application, not useful during operations.

**InfoLevel** - Normal operational messages - may be harvested for reporting, measuring throughput, etc. - no action required.

**NoticeLevel** - Normal but significant condition. Events that are unusual but not error conditions - might be summarized in an email to developers or admins to spot potential problems - no immediate action required.

**WarnLevel** - Warning messages, not an error, but indication that an error will occur if action is not taken, e.g. file system 85% full - each item must be resolved within a given time.

**ErrorLevel** - Non-urgent failures, these should be relayed to developers or admins; each item must be resolved within a given time.

**PanicLevel** - A "panic" condition usually affecting multiple apps/servers/sites. At this level it would usually notify all tech staff on call.

**AlertLevel** - Action must be taken immediately. Should be corrected immediately, therefore notify staff who can fix the problem. An example would be the loss of a primary ISP connection.

**FatalLevel** - Should be corrected immediately, but indicates failure in a primary system, an example is a loss of a backup ISP connection. ( same as SYSLOG CRITICAL )

## Contributing

We are welcome to any contributions. Just fork the
[project](https://github.com/phogolabs/log).

[report-img]: https://goreportcard.com/badge/github.com/phogolabs/log
[report-url]: https://goreportcard.com/report/github.com/phogolabs/log
[codecov-url]: https://codecov.io/gh/phogolabs/log
[codecov-img]: https://codecov.io/gh/phogolabs/log/branch/master/graph/badge.svg
[travis-img]: https://travis-ci.org/phogolabs/log.svg?branch=master
[travis-url]: https://travis-ci.org/phogolabs/log
[log-url]: https://github.com/phogolabs/prana
[godoc-url]: https://godoc.org/github.com/phogolabs/log
[godoc-img]: https://godoc.org/github.com/phogolabs/log?status.svg
[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[software-license-url]: LICENSE
