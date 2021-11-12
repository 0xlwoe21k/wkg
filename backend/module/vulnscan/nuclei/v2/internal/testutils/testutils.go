package testutils

import (
	catalog2 "backend/module/vulnscan/nuclei/v2/pkg/catalog"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	progress2 "backend/module/vulnscan/nuclei/v2/pkg/progress"
	protocols2 "backend/module/vulnscan/nuclei/v2/pkg/protocols"
	protocolinit2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/common/protocolinit"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"github.com/logrusorgru/aurora"
	"go.uber.org/ratelimit"

	"github.com/projectdiscovery/gologger/levels"
)

// Init initializes the protocols and their configurations
func Init(options *types2.Options) {
	_ = protocolinit2.Init(options)
}

// DefaultOptions is the default options structure for nuclei during mocking.
var DefaultOptions = &types2.Options{
	Metrics:            false,
	Debug:              false,
	DebugRequests:      false,
	DebugResponse:      false,
	Silent:             false,
	Version:            false,
	Verbose:            false,
	NoColor:            true,
	UpdateTemplates:    false,
	JSON:               false,
	JSONRequests:       false,
	EnableProgressBar:  false,
	TemplatesVersion:   false,
	TemplateList:       false,
	Stdin:              false,
	StopAtFirstMatch:   false,
	NoMeta:             false,
	Project:            false,
	MetricsPort:        0,
	BulkSize:           25,
	TemplateThreads:    10,
	Timeout:            5,
	Retries:            1,
	RateLimit:          150,
	ProjectPath:        "",
	Severities:         severity2.Severities{},
	Targets:            []string{},
	TargetsFilePath:    "",
	Output:             "",
	ProxyURL:           "",
	ProxySocksURL:      "",
	TemplatesDirectory: "",
	TraceLogFile:       "",
	Templates:          []string{},
	ExcludedTemplates:  []string{},
	CustomHeaders:      []string{},
}

// MockOutputWriter is a mocked output writer.
type MockOutputWriter struct {
	aurora          aurora.Aurora
	RequestCallback func(templateID, url, requestType string, err error)
	WriteCallback   func(o *output2.ResultEvent)
}

// NewMockOutputWriter creates a new mock output writer
func NewMockOutputWriter() *MockOutputWriter {
	return &MockOutputWriter{aurora: aurora.NewAurora(false)}
}

// Close closes the output writer interface
func (m *MockOutputWriter) Close() {}

// Colorizer returns the colorizer instance for writer
func (m *MockOutputWriter) Colorizer() aurora.Aurora {
	return m.aurora
}

// Write writes the event to file and/or screen.
func (m *MockOutputWriter) Write(result *output2.ResultEvent) error {
	if m.WriteCallback != nil {
		m.WriteCallback(result)
	}
	return nil
}

// Request writes a log the requests trace log
func (m *MockOutputWriter) Request(templateID, url, requestType string, err error) {
	if m.RequestCallback != nil {
		m.RequestCallback(templateID, url, requestType, err)
	}
}

// TemplateInfo contains info for a mock executed template.
type TemplateInfo struct {
	ID   string
	Info model2.Info
	Path string
}

// NewMockExecuterOptions creates a new mock executeroptions struct
func NewMockExecuterOptions(options *types2.Options, info *TemplateInfo) *protocols2.ExecuterOptions {
	progressImpl, _ := progress2.NewStatsTicker(0, false, false, false, 0)
	executerOpts := &protocols2.ExecuterOptions{
		TemplateID:   info.ID,
		TemplateInfo: info.Info,
		TemplatePath: info.Path,
		Output:       NewMockOutputWriter(),
		Options:      options,
		Progress:     progressImpl,
		ProjectFile:  nil,
		IssuesClient: nil,
		Browser:      nil,
		Catalog:      catalog2.New(options.TemplatesDirectory),
		RateLimiter:  ratelimit.New(options.RateLimit),
	}
	return executerOpts
}

// NoopWriter is a NooP gologger writer.
type NoopWriter struct{}

// Write writes the data to an output writer.
func (n *NoopWriter) Write(data []byte, level levels.Level) {}
