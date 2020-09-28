package appsflyer_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/export/qwerty1/qwerty?api_token=qwerty2", uri.String())
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
	assert.Equal(t, "https://github.com/export/qwerty1/qwerty?api_token=qwerty2&bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
}
