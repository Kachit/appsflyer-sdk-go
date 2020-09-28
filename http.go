package appsflyer_sdk

import (
	"fmt"
	"github.com/gocarina/gocsv"
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
	u.Path = "/" + path
	u.RawQuery = rb.buildQueryParams(query)
	return u, err
}

func (rb *RequestBuilder) buildQueryParams(query map[string]interface{}) string {
	q := url.Values{}
	if query != nil {
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return q.Encode()
}

type Response struct {
	raw *http.Response
}

func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

func (r *Response) GetRaw() *http.Response {
	return r.raw
}

func (r *Response) GetData() ([]*Report, error) {
	defer r.raw.Body.Close()
	body, err := ioutil.ReadAll(r.raw.Body)
	if err != nil {
		return nil, err
	}

	reports := []*Report{}
	err = gocsv.UnmarshalBytes(body, &reports)
	return reports, err
}

func (r *Response) UnmarshalCSV(reports []*Report) error {
	defer r.raw.Body.Close()
	body, err := ioutil.ReadAll(r.raw.Body)
	if err != nil {
		return err
	}
	return gocsv.UnmarshalBytes(body, &reports)
}
