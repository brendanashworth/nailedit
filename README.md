# nailedit

[![Build Status](https://travis-ci.org/brendanashworth/nailedit.svg?branch=master)](https://travis-ci.org/brendanashworth/nailedit)  [![codecov](https://codecov.io/gh/brendanashworth/nailedit/branch/master/graph/badge.svg)](https://codecov.io/gh/brendanashworth/nailedit) [![GoDoc](https://godoc.org/github.com/brendanashworth/nailedit?status.svg)](http://godoc.org/github.com/brendanashworth/nailedit/thumbnail)

> `nailedit` is a library that generates thumbnail images. It resizes
images and changes their format according to passed options and uses the
fastest available library to do so. It has both a command line interface
and a traditional [API](https://godoc.org/github.com/brendanashworth/nailedit/thumbnail).

## CLI Usage

`nailedit` takes input from STDIN and pipes the thumbnail to STDOUT.

```sh
$ cat image.jpeg | nailedit > thumbnail.jpeg

$ nailedit -help
```

## Options

- *`-dimensions`*: dimensions of the output thumbnail image, in format `WIDTHxHEIGHT`
- *`-format`*: format of the output thumbnail image (i.e., `jpeg`, `png`)
