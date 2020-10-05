# Appsflyer SDK GO (Unofficial)
[![Build Status](https://travis-ci.org/Kachit/appsflyer-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/appsflyer-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/appsflyer-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/appsflyer-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/appsflyer-sdk-go)](https://goreportcard.com/report/github.com/kachit/appsflyer-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/appsflyer-sdk-go.svg)](https://github.com/Kachit/appsflyer-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/appsflyer-sdk-go/blob/master/LICENSE)

## Description
Unofficial Appsflyer API Client for Golang (reporting)

## API documentation
https://support.appsflyer.com/hc/en-us/articles/209680773-Raw-data-reporting-overview

## Download
```shell
go get -u github.com/kachit/appsflyer-sdk-go
```

## Usage
```go
package main

import (
	"fmt"
	appsflyer_sdk "github.com/kachit/appsflyer-sdk-go"
	"time"
)

func main() {
    config := appsflyer_sdk.NewConfig("foo", "bar")
    client := appsflyer_sdk.NewClientFromConfig(config, nil)

    from := time.Date(2020, time.Month(9), 10, 0, 0, 0, 0, time.UTC)
    to := time.Date(2020, time.Month(9), 11, 0, 0, 0, 0, time.UTC)

    filter := &appsflyer_sdk.InstallsReportFilter{
        StartDate: from,
        EndDate: to,
    }
    response, err := client.Reports().GetInstallReports(filter)
    if err != nil {
        fmt.Println(err)
    }

    if !response.IsSuccess() {
        fmt.Println(response.GetError())
    }

    reports := []*appsflyer_sdk.Report{}
    err = response.UnmarshalCSV(reports)

    if err != nil {
        fmt.Println(err)
    }
    
    fmt.Println(reports)
}
```
