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
require github.com/disaster37/go-kibana-rest/v7
```

## Usage

### Init the client

```go
cfg := kibana.Config{
    Address:          "http://127.0.0.1:5601",
    Username:         "elastic",
    Password:         "changeme",
    DisableVerifySSL: true,
}

client, err := kibana.NewClient(cfg)

if err != nil {
    log.Fatalf("Error creating the client: %s", err)
}

status, err := client.API.KibanaStatus.Get()
if err != nil {
    log.Fatalf("Error getting response: %s", err)
}
log.Println(status)
```

### Handle shorten URL

```go

```

## Contribute

First, if you use kibana module that required license like Logstash Pipeline, you need to have valid license or start trial license.

Start trial license:
```bash
curl -XPOST -u elastic:changeme "http://localhost:9200/_license/start_trial?acknowledge=true&pretty"

```