# go-kibana-rest

[![CircleCI](https://circleci.com/gh/disaster37/go-kibana-rest.svg?style=svg)](https://circleci.com/gh/disaster37/go-kibana-rest)
[![Go Report Card](https://goreportcard.com/badge/github.com/disaster37/go-kibana-rest)](https://goreportcard.com/report/github.com/disaster37/go-kibana-rest)
[![GoDoc](https://godoc.org/github.com/disaster37/go-kibana-rest?status.svg)](http://godoc.org/github.com/disaster37/go-kibana-rest)
[![codecov](https://codecov.io/gh/disaster37/go-kibana-rest/branch/master/graph/badge.svg)](https://codecov.io/gh/disaster37/go-kibana-rest)

Go framework to handle kibana API

## Compatibility

At the moment is only work with Kibana 7.x.

## Installation

In your go.mod, put:
```go
require github.com/disaster37/go-kibana-rest
```

## Usage

@todo

## Contribute

First, if you use kibana module that required license like Logstash Pipeline, you need to have valid license or start trial license.

Start trial license:
```bash
curl -XPOST -u elastic:changeme "http://localhost:9200/_license/start_trial?acknowledge=true&pretty"

```