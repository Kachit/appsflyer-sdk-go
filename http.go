package appsflyer_sdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Transport struct {
	http *http.Client
	rb   *RequestBuilder
}

func newHttpTransport(config *Config, h *http.Client) *Transport {
	if h == nil {
		h = &http.Client{}
	}
	return &Transport{http: h, rb: &RequestBuilder{cfg: config}}
}

func (t *Transport) request(method string, path string, query map[string]interface{}, body map[string]interface{}) (resp *http.Response, err error) {
	//build uri
	uri, err := t.rb.buildUri(path, query)
	if err != nil {
		return nil, fmt.Errorf("transport@request build uri: %v", err)
	}
	//build request
	req, err := http.NewRequest(strings.ToUpper(method), uri.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("transport@request new request error: %v", err)
	}

	return t.http.Do(req)
}

func (t *Transport) get(path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.request("GET", path, query, nil)
}

type RequestBuilder struct {
	cfg *Config
}

func (rb *RequestBuilder) buildUri(path string, query map[string]interface{}) (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildUri parse: %v", err)
	}
	u.Path = rb.buildPath(path)
	u.RawQuery = rb.buildQueryParams(query)
	return u, err
}

func (rb *RequestBuilder) buildQueryParams(query map[string]interface{}) string {
	q := url.Values{}
	q.Set("api_token", rb.cfg.APIToken)
	if query != nil {
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return q.Encode()
}

func (rb *RequestBuilder) buildPath(path string) string {
	return "/export/" + rb.cfg.AppId + "/" + path
}

type Response struct {
	raw *http.Response
	csv *CSVParser
}

func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

func (r *Response) GetRaw() *http.Response {
	return r.raw
}

func (r *Response) GetData() ([]Report, error) {
	defer r.raw.Body.Close()
	body, err := ioutil.ReadAll(r.raw.Body)
	if err != nil {
		return nil, err
	}

	var entities []Report
	if err := r.csv.Parse(string(body), Report{}, func(v interface{}) {
		entities = append(entities, v.(Report))
	}); err != nil {
		return nil, err
	}
	return entities, nil
}
