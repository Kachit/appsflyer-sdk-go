package appsflyer_sdk

import (
	"strings"
	"time"
)

type Report struct {
	AttributedTouchType string `json:"attributed_touch_type" csv:"Attributed Touch Type"`
	AttributedTouchTime string `json:"attributed_touch_time" csv:"Attributed Touch Time"`
	InstallTime         string `json:"install_time" csv:"Install Time"`
	EventTime           string `json:"event_time" csv:"Event Time"`
	EventName           string `json:"event_name" csv:"Event Name"`
	MediaSource         string `json:"media_source" csv:"Media Source"`
	Channel             string `json:"channel" csv:"Channel"`
	Campaign            string `json:"campaign" csv:"Campaign"`
	Ad                  string `json:"ad" csv:"Ad"`
	AdvertisingID       string `json:"advertising_id" csv:"Advertising ID"`
	IDFA                string `json:"idfa" csv:"IDFA"`
	CustomerUserID      string `json:"customer_user_id" csv:"Customer User ID"`
	IsRetargeting       string `json:"is_retargeting" csv:"Is Retargeting"`
	IP                  string `json:"ip" csv:"IP"`
	AppsflyerID         string `json:"appsflyer_id" csv:"AppsFlyer ID"`
	AndroidID           string `json:"android_id" csv:"Android ID"`
	OSVersion           string `json:"os_version" csv:"OS Version"`
	AppVersion          string `json:"app_version" csv:"App Version"`
	SDKVersion          string `json:"sdk_version" csv:"SDK Version"`
	UserAgent           string `json:"user_agent" csv:"User Agent"`
	OriginalURL         string `json:"original_url" csv:"Original URL"`
	CountryCode         string `json:"country_code" csv:"Country Code"`
	Language            string `json:"language" csv:"Language"`
}

func (r *Report) GetIsRetargeting() bool {
	switch strings.ToLower(r.IsRetargeting) {
	case "false":
		return false
	case "true":
		return true
	}
	return false
}

type ReportsResource struct {
	*ResourceAbstract
}

func (rr *ReportsResource) GetInstallReports(filter *InstallsReportFilter) (*Response, error) {
	return rr.get("installs_report/v5", filter.Build())
}

type InstallsReportFilter struct {
	StartDate        time.Time
	EndDate          time.Time
	AdditionalFields []string
}

func (f *InstallsReportFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["from"] = f.StartDate.Format("2006-01-02")
	params["to"] = f.EndDate.Format("2006-01-02")
	params["timezone"] = f.StartDate.Location().String()
	if f.AdditionalFields != nil {
		params["additional_fields"] = strings.Join(f.AdditionalFields, ",")
	}
	return params
}

type AppsEventReportFilter struct {
	StartDate        time.Time
	EndDate          time.Time
	AdditionalFields []string
}

func (f *AppsEventReportFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["from"] = f.StartDate.Format("2006-01-02")
	params["to"] = f.EndDate.Format("2006-01-02")
	params["timezone"] = f.StartDate.Location().String()
	if f.AdditionalFields != nil {
		params["additional_fields"] = strings.Join(f.AdditionalFields, ",")
	}
	return params
}
