package appsflyer_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Reports_InstallsReportFilter_Build(t *testing.T) {
	filter := InstallsReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	expected := make(map[string]interface{})
	expected["from"] = "2020-01-10"
	expected["to"] = "2020-01-20"
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Reports_AppsEventReportFilter_Build(t *testing.T) {
	filter := AppsEventReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	expected := make(map[string]interface{})
	expected["from"] = "2020-01-10"
	expected["to"] = "2020-01-20"
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Reports_InstallsReportFilter_IsValidSuccess(t *testing.T) {
	filter := InstallsReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	assert.Nil(t, filter.IsValid())
}

func Test_Reports_InstallsReportFilter_IsValidFailedDateFrom(t *testing.T) {
	filter := InstallsReportFilter{}
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "InstallsReportFilter@IsValid: StartDate is required", err.Error())
}

func Test_Reports_InstallsReportFilter_IsValidFailedDateTo(t *testing.T) {
	filter := InstallsReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "InstallsReportFilter@IsValid: EndDate is required", err.Error())
}

func Test_Reports_AppsEventReportFilter_IsValidSuccess(t *testing.T) {
	filter := AppsEventReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	assert.Nil(t, filter.IsValid())
}

func Test_Reports_AppsEventReportFilter_IsValidFailedDateFrom(t *testing.T) {
	filter := AppsEventReportFilter{}
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "AppsEventReportFilter@IsValid: StartDate is required", err.Error())
}

func Test_Reports_AppsEventReportFilter_IsValidFailedDateTo(t *testing.T) {
	filter := AppsEventReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "AppsEventReportFilter@IsValid: EndDate is required", err.Error())
}
