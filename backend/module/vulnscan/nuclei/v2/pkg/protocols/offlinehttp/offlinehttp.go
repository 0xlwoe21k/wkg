package offlinehttp

import (
	operators2 "backend/module/vulnscan/nuclei/v2/pkg/operators"
	protocols2 "backend/module/vulnscan/nuclei/v2/pkg/protocols"
	"github.com/pkg/errors"
)

// Request is a offline http response processing request
type Request struct {
	options           *protocols2.ExecuterOptions
	compiledOperators []*operators2.Operators
}

// GetID returns the unique ID of the request if any.
func (r *Request) GetID() string {
	return ""
}

// Compile compiles the protocol request for further execution.
func (r *Request) Compile(options *protocols2.ExecuterOptions) error {
	for _, operator := range options.Operators {
		if err := operator.Compile(); err != nil {
			return errors.Wrap(err, "could not compile operators")
		}
		r.compiledOperators = append(r.compiledOperators, operator)
	}
	r.options = options
	return nil
}

// Requests returns the total number of requests the YAML rule will perform
func (r *Request) Requests() int {
	return 1
}
