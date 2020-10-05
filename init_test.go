package appsflyer_sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func buildStubConfig() *Config {
	cfg := &Config{
		Uri:      "https://github.com",
		AppId:    "qwerty1",
		APIToken: "qwerty2",
	}
	return cfg
}

func buildStubClient() *Client {
	return NewClientFromConfig(buildStubConfig(), nil)
}

func buildStubReportsResource() *ReportsResource {
	return buildStubClient().Reports()
}

func loadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func buildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

func buildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := loadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}
