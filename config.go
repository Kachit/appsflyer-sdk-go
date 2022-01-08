package appsflyer

const (
	ApiBaseUrl = "https://hq.appsflyer.com"
)

type Config struct {
	Uri      string
	APIToken string
	AppId    string
}

func NewConfig(apiToken string, appId string) *Config {
	return &Config{Uri: ApiBaseUrl, APIToken: apiToken, AppId: appId}
}
