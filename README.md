# Appsflyer SDK GO
[![Build Status](https://travis-ci.org/Kachit/mytarget-sdk-go.svg?branch=master)](https://travis-ci.org/Kachit/mytarget-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/mytarget-sdk-go/branch/master/graph/badge.svg)](https://codecov.io/gh/Kachit/mytarget-sdk-go)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/mytarget-sdk-go/blob/master/LICENSE)

## Description
Appsflyer API Client for Go

## API documentation


## Download
```shell
go get github.com/kachit/appsflyer-sdk-go
```

## Usage
```go
package main

import (
    "fmt"
    "net/http"
    "github.com/kachit/appsflyer_sdk"
)

func yourFuncName(){ 
    cfg := appsflyer_sdk.NewConfig()

    client := appsflyer_sdk.NewClient(cfg, &http.Client{})

    fmt.Print(client)
}

```
