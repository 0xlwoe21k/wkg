package dns

import (
	testutils2 "backend/module/vulnscan/nuclei/v2/internal/testutils"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	operators2 "backend/module/vulnscan/nuclei/v2/pkg/operators"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDNSExecuteWithResults(t *testing.T) {
	options := testutils2.DefaultOptions

	testutils2.Init(options)
	templateID := "testing-dns"
	request := &Request{
		Type:      "A",
		Class:     "INET",
		Retries:   5,
		ID:        templateID,
		Recursion: false,
		Name:      "{{FQDN}}",
		Operators: operators2.Operators{
			Matchers: []*matchers2.Matcher{{
				Name:  "test",
				Part:  "raw",
				Type:  "word",
				Words: []string{"93.184.216.34"},
			}},
			Extractors: []*extractors2.Extractor{{
				Part:  "raw",
				Type:  "regex",
				Regex: []string{"[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+"},
			}},
		},
	}
	executerOpts := testutils2.NewMockExecuterOptions(options, &testutils2.TemplateInfo{
		ID:   templateID,
		Info: model2.Info{SeverityHolder: severity2.Holder{Severity: severity2.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile dns request")

	var finalEvent *output2.InternalWrappedEvent
	t.Run("domain-valid", func(t *testing.T) {
		metadata := make(output2.InternalEvent)
		previous := make(output2.InternalEvent)
		err := request.ExecuteWithResults("example.com", metadata, previous, func(event *output2.InternalWrappedEvent) {
			finalEvent = event
		})
		require.Nil(t, err, "could not execute dns request")
	})
	require.NotNil(t, finalEvent, "could not get event output from request")
	require.Equal(t, 1, len(finalEvent.Results), "could not get correct number of results")
	require.Equal(t, "test", finalEvent.Results[0].MatcherName, "could not get correct matcher name of results")
	require.Equal(t, 1, len(finalEvent.Results[0].ExtractedResults), "could not get correct number of extracted results")
	require.Equal(t, "93.184.216.34", finalEvent.Results[0].ExtractedResults[0], "could not get correct extracted results")
	finalEvent = nil

	t.Run("url-to-domain", func(t *testing.T) {
		metadata := make(output2.InternalEvent)
		previous := make(output2.InternalEvent)
		err := request.ExecuteWithResults("https://example.com", metadata, previous, func(event *output2.InternalWrappedEvent) {
			finalEvent = event
		})
		require.Nil(t, err, "could not execute dns request")
	})
	require.NotNil(t, finalEvent, "could not get event output from request")
	require.Equal(t, 1, len(finalEvent.Results), "could not get correct number of results")
	require.Equal(t, "test", finalEvent.Results[0].MatcherName, "could not get correct matcher name of results")
	require.Equal(t, 1, len(finalEvent.Results[0].ExtractedResults), "could not get correct number of extracted results")
	require.Equal(t, "93.184.216.34", finalEvent.Results[0].ExtractedResults[0], "could not get correct extracted results")
	finalEvent = nil
}
