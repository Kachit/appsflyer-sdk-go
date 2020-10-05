package appsflyer_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	result := NewConfig("foo", "bar")
	assert.NotEmpty(t, result)
}
