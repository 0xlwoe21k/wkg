//go:generate dstdocgen -path "" -structure Template -output templates_doc.go -package templates
package templates

import (
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	protocols2 "backend/module/vulnscan/nuclei/v2/pkg/protocols"
	dns2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/dns"
	file2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/file"
	headless2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/headless"
	http2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/http"
	network2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/network"
	workflows2 "backend/module/vulnscan/nuclei/v2/pkg/workflows"
)

// Template is a YAML input file which defines all the requests and
// other metadata for a template.
type Template struct {
	// description: |
	//   ID is the unique id for the template.
	//
	//   #### Good IDs
	//
	//   A good ID uniquely identifies what the requests in the template
	//   are doing. Let's say you have a template that identifies a git-config
	//   file on the webservers, a good name would be `git-config-exposure`. Another
	//   example name is `azure-apps-nxdomain-takeover`.
	// examples:
	//   - name: ID Example
	//     value: "\"CVE-2021-19520\""
	ID string `yaml:"id" jsonschema:"title=id of the template,description=The Unique ID for the template,example=cve-2021-19520"`
	// description: |
	//   Info contains metadata information about the template.
	// examples:
	//   - value: exampleInfoStructure
	Info model2.Info `yaml:"info" jsonschema:"title=info for the template,description=Info contains metadata for the template"`
	// description: |
	//   Requests contains the http request to make in the template.
	// examples:
	//   - value: exampleNormalHTTPRequest
	RequestsHTTP []*http2.Request `yaml:"requests,omitempty" json:"requests,omitempty" jsonschema:"title=http requests to make,description=HTTP requests to make for the template"`
	// description: |
	//   DNS contains the dns request to make in the template
	// examples:
	//   - value: exampleNormalDNSRequest
	RequestsDNS []*dns2.Request `yaml:"dns,omitempty" json:"dns,omitempty" jsonschema:"title=dns requests to make,description=DNS requests to make for the template"`
	// description: |
	//   File contains the file request to make in the template
	// examples:
	//   - value: exampleNormalFileRequest
	RequestsFile []*file2.Request `yaml:"file,omitempty" json:"file,omitempty" jsonschema:"title=file requests to make,description=File requests to make for the template"`
	// description: |
	//   Network contains the network request to make in the template
	// examples:
	//   - value: exampleNormalNetworkRequest
	RequestsNetwork []*network2.Request `yaml:"network,omitempty" json:"network,omitempty" jsonschema:"title=network requests to make,description=Network requests to make for the template"`
	// description: |
	//   Headless contains the headless request to make in the template.
	RequestsHeadless []*headless2.Request `yaml:"headless,omitempty" json:"headless,omitempty" jsonschema:"title=headless requests to make,description=Headless requests to make for the template"`

	// description: |
	//   Workflows is a yaml based workflow declaration code.
	workflows2.Workflow `yaml:",inline,omitempty" jsonschema:"title=workflows to run,description=Workflows to run for the template"`
	CompiledWorkflow    *workflows2.Workflow `yaml:"-" json:"-" jsonschema:"-"`

	// TotalRequests is the total number of requests for the template.
	TotalRequests int `yaml:"-" json:"-"`
	// Executer is the actual template executor for running template requests
	Executer protocols2.Executer `yaml:"-" json:"-"`

	Path string `yaml:"-" json:"-"`
}
