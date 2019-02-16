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
