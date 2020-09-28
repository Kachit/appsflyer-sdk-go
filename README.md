# Appsflyer SDK GO
[![Build Status](https://travis-ci.org/Kachit/appsflyer-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/appsflyer-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/appsflyer-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/appsflyer-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/appsflyer-sdk-go/blob/master/LICENSE)

## Description
Appsflyer API Client for Go (reporting)

## API documentation
https://support.appsflyer.com/hc/en-us/articles/209680773-Raw-data-reporting-overview

## Download
```shell
go get github.com/kachit/appsflyer-sdk-go
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
	location, _ := time.LoadLocation("Europe/Moscow")
	
	from := time.Date(2020, time.Month(9), 10, 0, 0, 0, 0, location)
	to := time.Date(2020, time.Month(9), 11, 0, 0, 0, 0, location)
	
	config := appsflyer_sdk.NewConfig("foo", "bar")
	client := appsflyer_sdk.NewClient(config, nil)
	
	filter := &appsflyer_sdk.InstallsReportFilter{
		StartDate: from,
		EndDate: to,
	}
	response, err := client.Reports().GetInstallReports(filter)
    if err != nil {
        fmt.Println(response.GetData())
    }
	
	fmt.Println(err)
}
```
