package appsflyer_sdk

import "fmt"

type ResourceAbstract struct {
	tr  *Transport
	csv *CSVParser
}

func (r *ResourceAbstract) get(path string, query map[string]interface{}) (*Response, error) {
	rsp, err := r.tr.get(path, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@get request: %v", err)
	}
	return &Response{raw: rsp, csv: r.csv}, nil
}

func newResourceAbstract(transport *Transport) *ResourceAbstract {
	return &ResourceAbstract{tr: transport, csv: &CSVParser{}}
}
