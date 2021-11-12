package dns

import (
	testutils2 "backend/module/vulnscan/nuclei/v2/internal/testutils"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	operators2 "backend/module/vulnscan/nuclei/v2/pkg/operators"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	"net"
	"strconv"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/require"
)

func TestResponseToDSLMap(t *testing.T) {
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
	}
	executerOpts := testutils2.NewMockExecuterOptions(options, &testutils2.TemplateInfo{
		ID:   templateID,
		Info: model2.Info{SeverityHolder: severity2.Holder{Severity: severity2.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile dns request")

	req := new(dns.Msg)
	req.Question = append(req.Question, dns.Question{Name: "one.one.one.one.", Qtype: dns.TypeA, Qclass: dns.ClassINET})

	resp := new(dns.Msg)
	resp.Rcode = dns.RcodeSuccess
	resp.Answer = append(resp.Answer, &dns.A{A: net.ParseIP("1.1.1.1"), Hdr: dns.RR_Header{Name: "one.one.one.one."}})

	event := request.responseToDSLMap(req, resp, "one.one.one.one", "one.one.one.one")
	require.Len(t, event, 12, "could not get correct number of items in dsl map")
	require.Equal(t, dns.RcodeSuccess, event["rcode"], "could not get correct rcode")
}

func TestDNSOperatorMatch(t *testing.T) {
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
	}
	executerOpts := testutils2.NewMockExecuterOptions(options, &testutils2.TemplateInfo{
		ID:   templateID,
		Info: model2.Info{SeverityHolder: severity2.Holder{Severity: severity2.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile dns request")

	req := new(dns.Msg)
	req.Question = append(req.Question, dns.Question{Name: "one.one.one.one.", Qtype: dns.TypeA, Qclass: dns.ClassINET})

	resp := new(dns.Msg)
	resp.Rcode = dns.RcodeSuccess
	resp.Answer = append(resp.Answer, &dns.A{A: net.ParseIP("1.1.1.1"), Hdr: dns.RR_Header{Name: "one.one.one.one."}})

	event := request.responseToDSLMap(req, resp, "one.one.one.one", "one.one.one.one")

	t.Run("valid", func(t *testing.T) {
		matcher := &matchers2.Matcher{
			Part:  "raw",
			Type:  "word",
			Words: []string{"1.1.1.1"},
		}
		err = matcher.CompileMatchers()
		require.Nil(t, err, "could not compile matcher")

		matched := request.Match(event, matcher)
		require.True(t, matched, "could not match valid response")
	})

	t.Run("rcode", func(t *testing.T) {
		matcher := &matchers2.Matcher{
			Part:   "rcode",
			Type:   "status",
			Status: []int{dns.RcodeSuccess},
		}
		err = matcher.CompileMatchers()
		require.Nil(t, err, "could not compile rcode matcher")

		matched := request.Match(event, matcher)
		require.True(t, matched, "could not match valid rcode response")
	})

	t.Run("negative", func(t *testing.T) {
		matcher := &matchers2.Matcher{
			Part:     "raw",
			Type:     "word",
			Negative: true,
			Words:    []string{"random"},
		}
		err := matcher.CompileMatchers()
		require.Nil(t, err, "could not compile negative matcher")

		matched := request.Match(event, matcher)
		require.True(t, matched, "could not match valid negative response matcher")
	})

	t.Run("invalid", func(t *testing.T) {
		matcher := &matchers2.Matcher{
			Part:  "raw",
			Type:  "word",
			Words: []string{"random"},
		}
		err := matcher.CompileMatchers()
		require.Nil(t, err, "could not compile matcher")

		matched := request.Match(event, matcher)
		require.False(t, matched, "could match invalid response matcher")
	})
}

func TestDNSOperatorExtract(t *testing.T) {
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
	}
	executerOpts := testutils2.NewMockExecuterOptions(options, &testutils2.TemplateInfo{
		ID:   templateID,
		Info: model2.Info{SeverityHolder: severity2.Holder{Severity: severity2.Low}, Name: "test"},
	})
	err := request.Compile(executerOpts)
	require.Nil(t, err, "could not compile dns request")

	req := new(dns.Msg)
	req.Question = append(req.Question, dns.Question{Name: "one.one.one.one.", Qtype: dns.TypeA, Qclass: dns.ClassINET})

	resp := new(dns.Msg)
	resp.Rcode = dns.RcodeSuccess
	resp.Answer = append(resp.Answer, &dns.A{A: net.ParseIP("1.1.1.1"), Hdr: dns.RR_Header{Name: "one.one.one.one."}})

	event := request.responseToDSLMap(req, resp, "one.one.one.one", "one.one.one.one")

	t.Run("extract", func(t *testing.T) {
		extractor := &extractors2.Extractor{
			Part:  "raw",
			Type:  "regex",
			Regex: []string{"[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+"},
		}
		err = extractor.CompileExtractors()
		require.Nil(t, err, "could not compile extractor")

		data := request.Extract(event, extractor)
		require.Greater(t, len(data), 0, "could not extractor valid response")
		require.Equal(t, map[string]struct{}{"1.1.1.1": {}}, data, "could not extract correct data")
	})

	t.Run("kval", func(t *testing.T) {
		extractor := &extractors2.Extractor{
			Type: "kval",
			KVal: []string{"rcode"},
		}
		err = extractor.CompileExtractors()
		require.Nil(t, err, "could not compile kval extractor")

		data := request.Extract(event, extractor)
		require.Greater(t, len(data), 0, "could not extractor kval valid response")
		require.Equal(t, map[string]struct{}{strconv.Itoa(dns.RcodeSuccess): {}}, data, "could not extract correct kval data")
	})
}

func TestDNSMakeResult(t *testing.T) {
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
	require.Nil(t, err, "could not compile dns request")

	req := new(dns.Msg)
	req.Question = append(req.Question, dns.Question{Name: "one.one.one.one.", Qtype: dns.TypeA, Qclass: dns.ClassINET})

	resp := new(dns.Msg)
	resp.Rcode = dns.RcodeSuccess
	resp.Answer = append(resp.Answer, &dns.A{A: net.ParseIP("1.1.1.1"), Hdr: dns.RR_Header{Name: "one.one.one.one."}})

	event := request.responseToDSLMap(req, resp, "one.one.one.one", "one.one.one.one")
	finalEvent := &output2.InternalWrappedEvent{InternalEvent: event}
	if request.CompiledOperators != nil {
		result, ok := request.CompiledOperators.Execute(event, request.Match, request.Extract)
		if ok && result != nil {
			finalEvent.OperatorsResult = result
			finalEvent.Results = request.MakeResultEvent(finalEvent)
		}
	}
	require.Equal(t, 1, len(finalEvent.Results), "could not get correct number of results")
	require.Equal(t, "test", finalEvent.Results[0].MatcherName, "could not get correct matcher name of results")
	require.Equal(t, "1.1.1.1", finalEvent.Results[0].ExtractedResults[0], "could not get correct extracted results")
}
