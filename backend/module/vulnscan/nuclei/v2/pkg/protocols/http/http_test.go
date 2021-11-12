package http

import (
	testutils2 "backend/module/vulnscan/nuclei/v2/internal/testutils"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPCompile(t *testing.T) {
	options := testutils2.DefaultOptions
	options.CustomHeaders = []string{"User-Agent: test", "Hello: World"}

	testutils2.Init(options)
	templateID := "testing-http"
	request := &Request{
		Name: "testing",
		Payloads: map[string]interface{}{
			"username": []string{"admin"},
			"password": []string{"admin", "guest", "password", "test", "12345", "123456"},
		},
		AttackType: "clusterbomb",
		Raw: []string{`GET /manager/html HTTP/1.1
Host: {{Hostname}}
User-Agent: Nuclei - Open-source project (github.com/projectdiscovery/nuclei)
Connection: close
Authorization: Basic {{username + ':' + password}}
Accept-Encoding: gzip`},
	}
	executerOpts := testutils2.NewMockExecuterOptions(options, &testutils2.TemplateInfo{
		ID:   templateID,
		Info: model2.Info{SeverityHolder: severity2.Holder{Severity: severity2.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile http request")
	require.Equal(t, 6, request.Requests(), "could not get correct number of requests")
	require.Equal(t, map[string]string{"User-Agent": "test", "Hello": "World"}, request.customHeaders, "could not get correct custom headers")
}
