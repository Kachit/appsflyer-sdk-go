package appsflyer_sdk

import "net/http"

type Client struct {
	transport *Transport
	config    *Config
}

func (c *Client) Reports() *ReportsResource {
	resource := newResourceAbstract(c.transport, c.config)
	return NewReportsResource(resource)
}

func NewClientFromConfig(config *Config, cl *http.Client) *Client {
	if cl == nil {
		cl = &http.Client{}
	}
	transport := newHttpTransport(config, cl)
	return &Client{transport: transport, config: config}
}

func NewClientFromCredentials(apiToken string, appId string, cl *http.Client) *Client {
	if cl == nil {
		cl = &http.Client{}
	}
	config := NewConfig(apiToken, appId)
	transport := newHttpTransport(config, cl)
	return &Client{transport: transport, config: config}
}
