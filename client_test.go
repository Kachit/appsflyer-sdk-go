package appsflyer_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_ReportsResource(t *testing.T) {
	client := buildStubClient()
	result := client.Reports()
	assert.NotEmpty(t, result)
}
