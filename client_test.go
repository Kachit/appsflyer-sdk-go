package appsflyer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_ReportsResource(t *testing.T) {
	client := buildStubClient()
	result := client.Reports()
	assert.NotEmpty(t, result)
}

func Test_Client_NewClientFromConfig(t *testing.T) {
	config := buildStubConfig()
	result := NewClientFromConfig(config, nil)
	assert.NotEmpty(t, result)
}

func Test_Client_NewClientFromCredentials(t *testing.T) {
	result := NewClientFromCredentials("foo", "bar", nil)
	assert.NotEmpty(t, result)
}
