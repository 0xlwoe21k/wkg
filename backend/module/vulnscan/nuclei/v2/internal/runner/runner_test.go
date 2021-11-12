package runner

import (
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createReportingOptions(t *testing.T) {
	var options types2.Options
	options.ReportingConfig = "../../../integration_tests/test-issue-tracker-config1.yaml"
	resultOptions, err := createReportingOptions(&options)

	assert.Nil(t, err)
	assert.Equal(t, resultOptions.AllowList.Severities, severity2.Severities{severity2.High, severity2.Critical})
	assert.Equal(t, resultOptions.DenyList.Severities, severity2.Severities{severity2.Low})

	options.ReportingConfig = "../../../integration_tests/test-issue-tracker-config2.yaml"
	resultOptions2, err := createReportingOptions(&options)
	assert.Nil(t, err)
	assert.Equal(t, resultOptions2.AllowList.Severities, resultOptions.AllowList.Severities)
	assert.Equal(t, resultOptions2.DenyList.Severities, resultOptions.DenyList.Severities)
}
