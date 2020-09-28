package appsflyer_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Common_CustomTimestamp_UnmarshalCSVFilled(t *testing.T) {
	c := CustomTimestamp{}
	format := "2006-01-02 15:04:05"
	str := "2020-09-10 15:15:15"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, "2020-09-10 15:15:15", c.Value().Format(format))
}

func Test_Common_CustomTimestamp_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomTimestamp{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.True(t, c.Value().IsZero())
}

func Test_Common_CustomBoolean_UnmarshalCSVTrue(t *testing.T) {
	c := CustomBoolean{}
	str := "true"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, true, c.Value())
}

func Test_Common_CustomBoolean_UnmarshalCSVFalse(t *testing.T) {
	c := CustomBoolean{}
	str := "false"
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, false, c.Value())
}

func Test_Common_CustomBoolean_UnmarshalCSVEmpty(t *testing.T) {
	c := CustomBoolean{}
	str := ""
	_ = c.UnmarshalCSV(str)
	assert.Equal(t, false, c.Value())
}
