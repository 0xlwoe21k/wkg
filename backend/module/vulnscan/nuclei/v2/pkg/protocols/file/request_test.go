package file

import (
	testutils2 "backend/module/vulnscan/nuclei/v2/internal/testutils"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	operators2 "backend/module/vulnscan/nuclei/v2/pkg/operators"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileExecuteWithResults(t *testing.T) {
	options := testutils2.DefaultOptions

	testutils2.Init(options)
	templateID := "testing-file"
	request := &Request{
		ID:                templateID,
		MaxSize:           1024,
		NoRecursive:       false,
		Extensions:        []string{"all"},
		ExtensionDenylist: []string{".go"},
		Operators: operators2.Operators{
			Matchers: []*matchers2.Matcher{{
				Name:  "test",
				Part:  "raw",
				Type:  "word",
				Words: []string{"1.1.1.1"},
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
	require.Nil(t, err, "could not compile file request")

	tempDir, err := ioutil.TempDir("", "test-*")
	require.Nil(t, err, "could not create temporary directory")
	defer os.RemoveAll(tempDir)

	files := map[string]string{
		"config.yaml": "TEST\r\n1.1.1.1\r\n",
	}
	for k, v := range files {
		err = ioutil.WriteFile(filepath.Join(tempDir, k), []byte(v), 0777)
		require.Nil(t, err, "could not write temporary file")
	}

	var finalEvent *output2.InternalWrappedEvent
	t.Run("valid", func(t *testing.T) {
		metadata := make(output2.InternalEvent)
		previous := make(output2.InternalEvent)
		err := request.ExecuteWithResults(tempDir, metadata, previous, func(event *output2.InternalWrappedEvent) {
			finalEvent = event
		})
		require.Nil(t, err, "could not execute file request")
	})
	require.NotNil(t, finalEvent, "could not get event output from request")
	require.Equal(t, 1, len(finalEvent.Results), "could not get correct number of results")
	require.Equal(t, "test", finalEvent.Results[0].MatcherName, "could not get correct matcher name of results")
	require.Equal(t, 1, len(finalEvent.Results[0].ExtractedResults), "could not get correct number of extracted results")
	require.Equal(t, "1.1.1.1", finalEvent.Results[0].ExtractedResults[0], "could not get correct extracted results")
	finalEvent = nil
}
