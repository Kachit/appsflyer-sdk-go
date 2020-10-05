package appsflyer_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Reports_InstallsReportFilter_BuildRequired(t *testing.T) {
	filter := InstallsReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)

	expected := make(map[string]interface{})
	expected["from"] = "2020-01-10"
	expected["to"] = "2020-01-20"
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Reports_InstallsReportFilter_BuildFull(t *testing.T) {
	filter := InstallsReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.UseTimezone = true
	filter.AdditionalFields = []string{"foo"}

	expected := make(map[string]interface{})
	expected["from"] = "2020-01-10"
	expected["to"] = "2020-01-20"
	expected["additional_fields"] = "foo"
	expected["timezone"] = "UTC"
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Reports_AppsEventReportFilter_BuildRequired(t *testing.T) {
	filter := AppsEventReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	expected := make(map[string]interface{})
	expected["from"] = "2020-01-10"
	expected["to"] = "2020-01-20"
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Reports_AppsEventReportFilter_BuildFull(t *testing.T) {
	filter := AppsEventReportFilter{}
	filter.StartDate = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.UseTimezone = true
	filter.AdditionalFields = []string{"foo"}

	expected := make(map[string]interface{})
	expected["from"] = "2020-01-10"
	expected["to"] = "2020-01-20"
	expected["additional_fields"] = "foo"
	expected["timezone"] = "UTC"
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

func Test_Reports_ReportsRequestBuilder_BuildQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := ReportsRequestBuilder{cfg: cfg}
	params := make(map[string]interface{})
	expected := make(map[string]interface{})
	expected["api_token"] = cfg.APIToken
	result := builder.buildQueryParams(params)
	assert.Equal(t, expected, result)
}

func Test_Reports_ReportsRequestBuilder_BuildPath(t *testing.T) {
	cfg := buildStubConfig()
	builder := ReportsRequestBuilder{cfg: cfg}
	result := builder.buildPath("installs_report/v5")
	assert.Equal(t, "/export/qwerty1/installs_report/v5", result)
}

func Test_Reports_ReportsResource_GetInstallReportsInvalidStartDate(t *testing.T) {
	filter := &InstallsReportFilter{}
	client := buildStubClient()
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	_, err := client.Reports().GetInstallReports(filter)
	assert.Error(t, err)
	assert.Equal(t, "ReportsResource@GetInstallReports: InstallsReportFilter@IsValid: StartDate is required", err.Error())
}

func Test_Reports_ReportsResource_GetInstallReportsInvalidEndDate(t *testing.T) {
	filter := &InstallsReportFilter{}
	client := buildStubClient()
	filter.StartDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	_, err := client.Reports().GetInstallReports(filter)
	assert.Error(t, err)
	assert.Equal(t, "ReportsResource@GetInstallReports: InstallsReportFilter@IsValid: EndDate is required", err.Error())
}

func Test_Reports_ReportsResource_GetAppsEventReportsInvalidStartDate(t *testing.T) {
	filter := &AppsEventReportFilter{}
	client := buildStubClient()
	filter.EndDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	_, err := client.Reports().GetAppsEventReports(filter)
	assert.Error(t, err)
	assert.Equal(t, "ReportsResource@GetAppsEventReports: AppsEventReportFilter@IsValid: StartDate is required", err.Error())
}

func Test_Reports_ReportsResource_GetAppsEventReportsInvalidEndDate(t *testing.T) {
	filter := &AppsEventReportFilter{}
	client := buildStubClient()
	filter.StartDate = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	_, err := client.Reports().GetAppsEventReports(filter)
	assert.Error(t, err)
	assert.Equal(t, "ReportsResource@GetAppsEventReports: AppsEventReportFilter@IsValid: EndDate is required", err.Error())
}
