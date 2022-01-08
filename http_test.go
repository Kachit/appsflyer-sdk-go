package appsflyer

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_Transport_Request(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := newHttpTransport(cfg, nil)

	body, _ := loadStubResponseData("stubs/data/reports/installs.csv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.request("GET", "foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGET(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := newHttpTransport(cfg, nil)

	body, _ := loadStubResponseData("stubs/data/reports/installs.csv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.get("foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: buildStubResponseFromFile(http.StatusOK, "stubs/data/reports/installs.csv")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: buildStubResponseFromFile(http.StatusBadRequest, "stubs/data/reports/installs.csv")}
	assert.False(t, response.IsSuccess())
}

func Test_HTTP_Response_GetRawResponse(t *testing.T) {
	rsp := buildStubResponseFromFile(http.StatusOK, "stubs/data/reports/installs.csv")
	response := &Response{raw: rsp}
	raw := response.GetRawResponse()
	assert.NotEmpty(t, raw)
	assert.Equal(t, http.StatusOK, raw.StatusCode)
}

func Test_HTTP_Response_GetRawBody(t *testing.T) {
	data, _ := loadStubResponseData("stubs/data/reports/installs.csv")
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/data/reports/installs.csv")
	response := &Response{raw: rsp}
	str, _ := response.GetRawBody()
	assert.Equal(t, string(data), str)
}
