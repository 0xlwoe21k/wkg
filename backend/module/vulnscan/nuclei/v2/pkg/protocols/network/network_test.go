package network

import (
	testutils2 "backend/module/vulnscan/nuclei/v2/internal/testutils"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNetworkCompileMake(t *testing.T) {
	options := testutils2.DefaultOptions

	testutils2.Init(options)
	templateID := "testing-network"
	request := &Request{
		ID:       templateID,
		Address:  []string{"{{Hostname}}", "{{Hostname}}:8082", "tls://{{Hostname}}:443"},
		ReadSize: 1024,
		Inputs:   []*Input{{Data: "test-data"}},
	}
	executerOpts := testutils2.NewMockExecuterOptions(options, &testutils2.TemplateInfo{
		ID:   templateID,
		Info: model2.Info{SeverityHolder: severity2.Holder{Severity: severity2.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile network request")

	require.Equal(t, 3, len(request.addresses), "could not get correct number of input address")
	t.Run("check-host", func(t *testing.T) {
		require.Equal(t, "{{Hostname}}", request.addresses[0].ip, "could not get correct host")
	})
	t.Run("check-host-with-port", func(t *testing.T) {
		require.Equal(t, "{{Hostname}}", request.addresses[1].ip, "could not get correct host with port")
		require.Equal(t, "8082", request.addresses[1].port, "could not get correct port for host")
	})
	t.Run("check-tls-with-port", func(t *testing.T) {
		require.Equal(t, "{{Hostname}}", request.addresses[2].ip, "could not get correct host with port")
		require.Equal(t, "443", request.addresses[2].port, "could not get correct port for host")
		require.True(t, request.addresses[2].tls, "could not get correct port for host")
	})
}
