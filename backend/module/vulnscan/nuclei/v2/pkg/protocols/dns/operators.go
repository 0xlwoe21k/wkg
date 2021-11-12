package dns

import (
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"bytes"
	"time"

	"github.com/miekg/dns"
)

// Match matches a generic data response again a given matcher
func (r *Request) Match(data map[string]interface{}, matcher *matchers2.Matcher) bool {
	partString := matcher.Part
	switch partString {
	case "body", "all", "":
		partString = "raw"
	}

	item, ok := data[partString]
	if !ok {
		return false
	}

	switch matcher.GetType() {
	case matchers2.StatusMatcher:
		return matcher.Result(matcher.MatchStatusCode(item.(int)))
	case matchers2.SizeMatcher:
		return matcher.Result(matcher.MatchSize(len(types2.ToString(item))))
	case matchers2.WordsMatcher:
		return matcher.Result(matcher.MatchWords(types2.ToString(item)))
	case matchers2.RegexMatcher:
		return matcher.Result(matcher.MatchRegex(types2.ToString(item)))
	case matchers2.BinaryMatcher:
		return matcher.Result(matcher.MatchBinary(types2.ToString(item)))
	case matchers2.DSLMatcher:
		return matcher.Result(matcher.MatchDSL(data))
	}
	return false
}

// Extract performs extracting operation for an extractor on model and returns true or false.
func (r *Request) Extract(data map[string]interface{}, extractor *extractors2.Extractor) map[string]struct{} {
	part := extractor.Part
	switch part {
	case "body", "all":
		part = "raw"
	}

	item, ok := data[part]
	if !ok {
		return nil
	}
	itemStr := types2.ToString(item)

	switch extractor.GetType() {
	case extractors2.RegexExtractor:
		return extractor.ExtractRegex(itemStr)
	case extractors2.KValExtractor:
		return extractor.ExtractKval(data)
	}
	return nil
}

// responseToDSLMap converts a DNS response to a map for use in DSL matching
func (r *Request) responseToDSLMap(req, resp *dns.Msg, host, matched string) output2.InternalEvent {
	data := make(output2.InternalEvent, 11)

	// Some data regarding the request metadata
	data["host"] = host
	data["matched"] = matched
	data["request"] = req.String()

	data["rcode"] = resp.Rcode
	buffer := &bytes.Buffer{}
	for _, question := range resp.Question {
		buffer.WriteString(question.String())
	}
	data["question"] = buffer.String()
	buffer.Reset()

	for _, extra := range resp.Extra {
		buffer.WriteString(extra.String())
	}
	data["extra"] = buffer.String()
	buffer.Reset()

	for _, answer := range resp.Answer {
		buffer.WriteString(answer.String())
	}
	data["answer"] = buffer.String()
	buffer.Reset()

	for _, ns := range resp.Ns {
		buffer.WriteString(ns.String())
	}
	data["ns"] = buffer.String()
	buffer.Reset()

	rawData := resp.String()
	data["raw"] = rawData
	data["template-id"] = r.options.TemplateID
	data["template-info"] = r.options.TemplateInfo
	data["template-path"] = r.options.TemplatePath
	return data
}

// MakeResultEvent creates a result event from internal wrapped event
func (r *Request) MakeResultEvent(wrapped *output2.InternalWrappedEvent) []*output2.ResultEvent {
	if len(wrapped.OperatorsResult.DynamicValues) > 0 {
		return nil
	}
	results := make([]*output2.ResultEvent, 0, len(wrapped.OperatorsResult.Matches)+1)

	// If we have multiple matchers with names, write each of them separately.
	if len(wrapped.OperatorsResult.Matches) > 0 {
		for k := range wrapped.OperatorsResult.Matches {
			data := r.makeResultEventItem(wrapped)
			data.MatcherName = k
			results = append(results, data)
		}
	} else if len(wrapped.OperatorsResult.Extracts) > 0 {
		for k, v := range wrapped.OperatorsResult.Extracts {
			data := r.makeResultEventItem(wrapped)
			data.ExtractedResults = v
			data.ExtractorName = k
			results = append(results, data)
		}
	} else {
		data := r.makeResultEventItem(wrapped)
		results = append(results, data)
	}
	return results
}

func (r *Request) makeResultEventItem(wrapped *output2.InternalWrappedEvent) *output2.ResultEvent {
	data := &output2.ResultEvent{
		TemplateID:       types2.ToString(wrapped.InternalEvent["template-id"]),
		TemplatePath:     types2.ToString(wrapped.InternalEvent["template-path"]),
		Info:             wrapped.InternalEvent["template-info"].(model2.Info),
		Type:             "dns",
		Host:             types2.ToString(wrapped.InternalEvent["host"]),
		Matched:          types2.ToString(wrapped.InternalEvent["matched"]),
		ExtractedResults: wrapped.OperatorsResult.OutputExtracts,
		Timestamp:        time.Now(),
		Request:          types2.ToString(wrapped.InternalEvent["request"]),
		Response:         types2.ToString(wrapped.InternalEvent["raw"]),
	}
	return data
}
