package file

import (
	testutils2 "backend/module/vulnscan/nuclei/v2/internal/testutils"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileCompile(t *testing.T) {
	options := testutils2.DefaultOptions

	testutils2.Init(options)
	templateID := "testing-file"
	request := &Request{
		ID:                templateID,
		MaxSize:           1024,
		NoRecursive:       false,
		Extensions:        []string{"all", ".lock"},
		ExtensionDenylist: []string{".go"},
	}
	executerOpts := testutils2.NewMockExecuterOptions(options, &testutils2.TemplateInfo{
		ID:   templateID,
		Info: model2.Info{SeverityHolder: severity2.Holder{Severity: severity2.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile file request")

	require.Contains(t, request.extensionDenylist, ".go", "could not get .go in denylist")
	require.NotContains(t, request.extensions, ".go", "could get .go in allowlist")
	require.True(t, request.allExtensions, "could not get correct allExtensions")
}
