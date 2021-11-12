package format

import (
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	stringslice2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/stringslice"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToMarkdownTableString(t *testing.T) {
	info := model2.Info{
		Name:           "Test Template Name",
		Authors:        stringslice2.StringSlice{Value: []string{"forgedhallpass", "ice3man"}},
		Description:    "Test description",
		SeverityHolder: severity2.Holder{Severity: severity2.High},
		Tags:           stringslice2.StringSlice{Value: []string{"cve", "misc"}},
		Reference:      stringslice2.StringSlice{Value: "reference1"},
		Metadata: map[string]string{
			"customDynamicKey1": "customDynamicValue1",
			"customDynamicKey2": "customDynamicValue2",
		},
	}

	result := ToMarkdownTableString(&info)

	expectedOrderedAttributes := `| Name | Test Template Name |
| Authors | forgedhallpass, ice3man |
| Tags | cve, misc |
| Severity | high |
| Description | Test description |`

	expectedDynamicAttributes := []string{
		"| customDynamicKey1 | customDynamicValue1 |",
		"| customDynamicKey2 | customDynamicValue2 |",
		"", // the expected result ends in a new line (\n)
	}

	actualAttributeSlice := strings.Split(result, "\n")
	dynamicAttributeIndex := len(actualAttributeSlice) - len(expectedDynamicAttributes)
	assert.Equal(t, strings.Split(expectedOrderedAttributes, "\n"), actualAttributeSlice[:dynamicAttributeIndex]) // the first part of the result is ordered
	assert.ElementsMatch(t, expectedDynamicAttributes, actualAttributeSlice[dynamicAttributeIndex:])              // dynamic parameters are not ordered
}
