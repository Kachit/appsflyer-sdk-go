package appsflyer_sdk

import (
	"fmt"
	"strings"
	"time"
)

type ReportsCollection struct {
	Reports []*Report
}

type Report struct {
	AttributedTouchType  string          `json:"attributed_touch_type" csv:"Attributed Touch Type"`
	AttributedTouchTime  CustomTimestamp `json:"attributed_touch_time" csv:"Attributed Touch Time"`
	InstallTime          CustomTimestamp `json:"install_time" csv:"Install Time"`
	EventTime            CustomTimestamp `json:"event_time" csv:"Event Time"`
	EventName            string          `json:"event_name" csv:"Event Name"`
	EventValue           string          `json:"event_value" csv:"Event Value"`
	EventRevenue         string          `json:"event_revenue" csv:"Event Revenue"`
	EventRevenueCurrency string          `json:"event_revenue_currency" csv:"Event Revenue Currency"`
	EventRevenueUSD      string          `json:"event_revenue_usd" csv:"Event Revenue USD"`
	EventSource          string          `json:"event_source" csv:"Event Source"`
	IsReceiptValidated   string          `json:"is_receipt_validated" csv:"Is Receipt Validated"`
	Partner              string          `json:"partner" csv:"Partner"`
	MediaSource          string          `json:"media_source" csv:"Media Source"`
	Channel              string          `json:"channel" csv:"Channel"`
	Keywords             string          `json:"keywords" csv:"Keywords"`
	Campaign             string          `json:"campaign" csv:"Campaign"`
	AdSet                string          `json:"adset" csv:"Adset"`
	AdsetID              string          `json:"adset_id" csv:"Adset ID"`
	Ad                   string          `json:"ad" csv:"Ad"`
	AdType               string          `json:"ad_type" csv:"Ad Type"`
	Region               string          `json:"region" csv:"Region"`
	CountryCode          string          `json:"country_code" csv:"Country Code"`
	State                string          `json:"state" csv:"State"`
	City                 string          `json:"city" csv:"City"`
	PostalCode           string          `json:"postal_code" csv:"Postal Code"`
	DMA                  string          `json:"dma" csv:"DMA"`
	IP                   string          `json:"ip" csv:"IP"`
	WIFI                 string          `json:"wifi" csv:"WIFI"`
	Language             string          `json:"language" csv:"Language"`
	AppsflyerID          string          `json:"appsflyer_id" csv:"AppsFlyer ID"`
	AdvertisingID        string          `json:"advertising_id" csv:"Advertising ID"`
	IDFA                 string          `json:"idfa" csv:"IDFA"`
	IDFV                 string          `json:"idfv" csv:"IDFV"`
	Platform             string          `json:"platform" csv:"Platform"`
	DeviceType           string          `json:"device_type" csv:"Device Type"`
	OSVersion            string          `json:"os_version" csv:"OS Version"`
	AppVersion           string          `json:"app_version" csv:"App Version"`
	SDKVersion           string          `json:"sdk_version" csv:"SDK Version"`
	AppID                string          `json:"app_id" csv:"App ID"`
	AppName              string          `json:"app_name" csv:"App Name"`
	BundleID             string          `json:"bundle_id" csv:"Bundle ID"`
	IsRetargeting        string          `json:"is_retargeting" csv:"Is Retargeting"`
	CustomerUserID       string          `json:"customer_user_id" csv:"Customer User ID"`

	AndroidID      string `json:"android_id" csv:"Android ID"`
	UserAgent      string `json:"user_agent" csv:"User Agent"`
	HTTPReferrer   string `json:"http_referrer" csv:"HTTP Referrer"`
	OriginalURL    string `json:"original_url" csv:"Original URL"`
	KeywordID      string `json:"keyword_id" csv:"Keyword ID"`
	StoreReinstall string `json:"store_reinstall" csv:"Store Reinstall"`
	DeeplinkURL    string `json:"deeplink_url" csv:"Deeplink URL"`

	OAID            string `json:"oaid" csv:"OAID"`
	InstallAppStore string `json:"install_app_store" csv:"Install App Store"`
	MatchType       string `json:"match_type" csv:"Match Type"`
	DeviceCategory  string `json:"device_category" csv:"Device Category"`
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
	UseTimezone      bool
	AdditionalFields []string
}

func (f *InstallsReportFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["from"] = f.StartDate.Format("2006-01-02")
	params["to"] = f.EndDate.Format("2006-01-02")
	if f.UseTimezone {
		params["timezone"] = f.StartDate.Location().String()
	}
	if f.AdditionalFields != nil {
		params["additional_fields"] = strings.Join(f.AdditionalFields, ",")
	}
	return params
}

func (f *InstallsReportFilter) IsValid() error {
	if f.StartDate.IsZero() {
		return fmt.Errorf("InstallsReportFilter@IsValid: %v", "StartDate is required")
	}
	if f.EndDate.IsZero() {
		return fmt.Errorf("InstallsReportFilter@IsValid: %v", "EndDate is required")
	}
	return nil
}

type AppsEventReportFilter struct {
	StartDate        time.Time
	EndDate          time.Time
	UseTimezone      bool
	AdditionalFields []string
}

func (f *AppsEventReportFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["from"] = f.StartDate.Format("2006-01-02")
	params["to"] = f.EndDate.Format("2006-01-02")
	if f.UseTimezone {
		params["timezone"] = f.StartDate.Location().String()
	}
	if f.AdditionalFields != nil {
		params["additional_fields"] = strings.Join(f.AdditionalFields, ",")
	}
	return params
}

func (f *AppsEventReportFilter) IsValid() error {
	if f.StartDate.IsZero() {
		return fmt.Errorf("AppsEventReportFilter@IsValid: %v", "StartDate is required")
	}
	if f.EndDate.IsZero() {
		return fmt.Errorf("AppsEventReportFilter@IsValid: %v", "EndDate is required")
	}
	return nil
}
