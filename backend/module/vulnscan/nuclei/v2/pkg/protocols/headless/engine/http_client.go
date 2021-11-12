package engine

import (
	protocolstate2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/common/protocolstate"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"crypto/tls"
	"net/http"
	"time"
)

// newhttpClient creates a new http client for headless communication with a timeout
func newhttpClient(options *types2.Options) *http.Client {
	dialer := protocolstate2.Dialer
	transport := &http.Transport{
		DialContext:         dialer.Dial,
		MaxIdleConns:        500,
		MaxIdleConnsPerHost: 500,
		MaxConnsPerHost:     500,
		TLSClientConfig: &tls.Config{
			Renegotiation:      tls.RenegotiateOnceAsClient,
			InsecureSkipVerify: true,
		},
	}
	return &http.Client{Transport: transport, Timeout: time.Duration(options.Timeout*3) * time.Second}
}
