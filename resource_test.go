package appsflyer

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Resources_Resource_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := newHttpTransport(cfg, nil)
	resource := newResourceAbstract(transport, cfg)

	body, _ := loadStubResponseData("stubs/data/reports/installs.csv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := resource.get("foo", nil)
	assert.NotEmpty(t, resp)
}
