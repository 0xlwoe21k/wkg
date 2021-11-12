package protocols

import (
	catalog2 "backend/module/vulnscan/nuclei/v2/pkg/catalog"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	operators2 "backend/module/vulnscan/nuclei/v2/pkg/operators"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	progress2 "backend/module/vulnscan/nuclei/v2/pkg/progress"
	projectfile2 "backend/module/vulnscan/nuclei/v2/pkg/projectfile"
	hosterrorscache2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/common/hosterrorscache"
	interactsh2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/common/interactsh"
	engine2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/headless/engine"
	reporting2 "backend/module/vulnscan/nuclei/v2/pkg/reporting"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"go.uber.org/ratelimit"
)

// Executer is an interface implemented any protocol based request executer.
type Executer interface {
	// Compile compiles the execution generators preparing any requests possible.
	Compile() error
	// Requests returns the total number of requests the rule will perform
	Requests() int
	// Execute executes the protocol group and returns true or false if results were found.
	Execute(input string) (bool, error)
	// ExecuteWithResults executes the protocol requests and returns results instead of writing them.
	ExecuteWithResults(input string, callback OutputEventCallback) error
}

// ExecuterOptions contains the configuration options for executer clients
type ExecuterOptions struct {
	// TemplateID is the ID of the template for the request
	TemplateID string
	// TemplatePath is the path of the template for the request
	TemplatePath string
	// TemplateInfo contains information block of the template request
	TemplateInfo model2.Info
	// Output is a writer interface for writing output events from executer.
	Output output2.Writer
	// Options contains configuration options for the executer.
	Options *types2.Options
	// IssuesClient is a client for nuclei issue tracker reporting
	IssuesClient *reporting2.Client
	// Progress is a progress client for scan reporting
	Progress progress2.Progress
	// RateLimiter is a rate-limiter for limiting sent number of requests.
	RateLimiter ratelimit.Limiter
	// Catalog is a template catalog implementation for nuclei
	Catalog *catalog2.Catalog
	// ProjectFile is the project file for nuclei
	ProjectFile *projectfile2.ProjectFile
	// Browser is a browser engine for running headless templates
	Browser *engine2.Browser
	// Interactsh is a client for interactsh oob polling server
	Interactsh *interactsh2.Client
	// HostErrorsCache is an optional cache for handling host errors
	HostErrorsCache *hosterrorscache2.Cache

	Operators []*operators2.Operators // only used by offlinehttp libs

	WorkflowLoader model2.WorkflowLoader
}

// Request is an interface implemented any protocol based request generator.
type Request interface {
	// Compile compiles the request generators preparing any requests possible.
	Compile(options *ExecuterOptions) error
	// Requests returns the total number of requests the rule will perform
	Requests() int
	// GetID returns the ID for the request if any. IDs are used for multi-request
	// condition matching. So, two requests can be sent and their match can
	// be evaluated from the third request by using the IDs for both requests.
	GetID() string
	// Match performs matching operation for a matcher on model and returns true or false.
	Match(data map[string]interface{}, matcher *matchers2.Matcher) bool
	// Extract performs extracting operation for an extractor on model and returns true or false.
	Extract(data map[string]interface{}, matcher *extractors2.Extractor) map[string]struct{}
	// ExecuteWithResults executes the protocol requests and returns results instead of writing them.
	ExecuteWithResults(input string, dynamicValues, previous output2.InternalEvent, callback OutputEventCallback) error
}

// OutputEventCallback is a callback event for any results found during scanning.
type OutputEventCallback func(result *output2.InternalWrappedEvent)
