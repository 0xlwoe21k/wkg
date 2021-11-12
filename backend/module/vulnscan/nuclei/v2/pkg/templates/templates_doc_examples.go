// Package templates
//nolint //do not lint as examples with no usage
package templates

import (
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	stringslice2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/stringslice"
	operators2 "backend/module/vulnscan/nuclei/v2/pkg/operators"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	dns2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/dns"
	file2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/file"
	http2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/http"
	network2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/network"
)

var (
	exampleInfoStructure = model2.Info{
		Name:           "Argument Injection in Ruby Dragonfly",
		Authors:        stringslice2.StringSlice{Value: "0xspara"},
		SeverityHolder: severity2.Holder{Severity: severity2.High},
		Reference:      stringslice2.StringSlice{Value: "https://zxsecurity.co.nz/research/argunment-injection-ruby-dragonfly/"},
		Tags:           stringslice2.StringSlice{Value: "cve,cve2021,rce,ruby"},
	}
	exampleNormalHTTPRequest = &http2.Request{
		Method: "GET",
		Path:   []string{"{{BaseURL}}/.git/config"},
		Operators: operators2.Operators{
			MatchersCondition: "and",
			Matchers: []*matchers2.Matcher{
				{Type: "word", Words: []string{"[core]"}},
				{Type: "dsl", DSL: []string{"!contains(tolower(body), '<html')", "!contains(tolower(body), '<body')"}, Condition: "and"},
				{Type: "status", Status: []int{200}}},
		},
	}
	_ = exampleNormalHTTPRequest

	exampleNormalDNSRequest = &dns2.Request{
		Name:      "{{FQDN}}",
		Type:      "CNAME",
		Class:     "inet",
		Retries:   2,
		Recursion: true,
		Operators: operators2.Operators{
			Extractors: []*extractors2.Extractor{
				{Type: "regex", Regex: []string{"ec2-[-\\d]+\\.compute[-\\d]*\\.amazonaws\\.com", "ec2-[-\\d]+\\.[\\w\\d\\-]+\\.compute[-\\d]*\\.amazonaws\\.com"}},
			},
		},
	}
	_ = exampleNormalDNSRequest

	exampleNormalFileRequest = &file2.Request{
		Extensions: []string{"all"},
		Operators: operators2.Operators{
			Extractors: []*extractors2.Extractor{
				{Type: "regex", Regex: []string{"amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"}},
			},
		},
	}
	_ = exampleNormalFileRequest

	exampleNormalNetworkRequest = &network2.Request{
		Inputs:   []*network2.Input{{Data: "envi\r\nquit\r\n"}},
		Address:  []string{"{{Hostname}}", "{{Hostname}}:2181"},
		ReadSize: 2048,
		Operators: operators2.Operators{
			Matchers: []*matchers2.Matcher{
				{Type: "word", Words: []string{"zookeeper.version"}},
			},
		},
	}
	_ = exampleNormalNetworkRequest
)
