package http

import (
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"net/http"
	"strings"
	"time"
)

// Match matches a generic data response again a given matcher
func (r *Request) Match(data map[string]interface{}, matcher *matchers2.Matcher) bool {
	item, ok := getMatchPart(matcher.Part, data)
	if !ok {
		return false
	}

	switch matcher.GetType() {
	case matchers2.StatusMatcher:
		statusCode, ok := data["status_code"]
		if !ok {
			return false
		}
		status, ok := statusCode.(int)
		if !ok {
			return false
		}
		return matcher.Result(matcher.MatchStatusCode(status))
	case matchers2.SizeMatcher:
		return matcher.Result(matcher.MatchSize(len(item)))
	case matchers2.WordsMatcher:
		return matcher.Result(matcher.MatchWords(item))
	case matchers2.RegexMatcher:
		return matcher.Result(matcher.MatchRegex(item))
	case matchers2.BinaryMatcher:
		return matcher.Result(matcher.MatchBinary(item))
	case matchers2.DSLMatcher:
		return matcher.Result(matcher.MatchDSL(data))
	}
	return false
}

// Extract performs extracting operation for an extractor on model and returns true or false.
func (r *Request) Extract(data map[string]interface{}, extractor *extractors2.Extractor) map[string]struct{} {
	item, ok := getMatchPart(extractor.Part, data)
	if !ok {
		return nil
	}
	switch extractor.GetType() {
	case extractors2.RegexExtractor:
		return extractor.ExtractRegex(item)
	case extractors2.KValExtractor:
		return extractor.ExtractKval(data)
	case extractors2.XPathExtractor:
		return extractor.ExtractHTML(item)
	case extractors2.JSONExtractor:
		return extractor.ExtractJSON(item)
	}
	return nil
}

// getMatchPart returns the match part honoring "all" matchers + others.
func getMatchPart(part string, data output2.InternalEvent) (string, bool) {
	if part == "header" {
		part = "all_headers"
	}
	var itemStr string

	if part == "all" {
		builder := &strings.Builder{}
		builder.WriteString(types2.ToString(data["body"]))
		builder.WriteString(types2.ToString(data["all_headers"]))
		itemStr = builder.String()
	} else {
		item, ok := data[part]
		if !ok {
			return "", false
		}
		itemStr = types2.ToString(item)
	}
	return itemStr, true
}

// responseToDSLMap converts an HTTP response to a map for use in DSL matching
func (r *Request) responseToDSLMap(resp *http.Response, host, matched, rawReq, rawResp, body, headers string, duration time.Duration, extra map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{}, len(extra)+8+len(resp.Header)+len(resp.Cookies()))
	for k, v := range extra {
		data[k] = v
	}
	for _, cookie := range resp.Cookies() {
		data[strings.ToLower(cookie.Name)] = cookie.Value
	}
	for k, v := range resp.Header {
		k = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(k), "-", "_"))
		data[k] = strings.Join(v, " ")
	}
	data["host"] = host
	data["matched"] = matched
	data["request"] = rawReq
	data["response"] = rawResp
	data["status_code"] = resp.StatusCode
	data["body"] = body
	data["content_length"] = resp.ContentLength
	data["all_headers"] = headers
	data["duration"] = duration.Seconds()
	data["template-id"] = r.options.TemplateID
	data["template-info"] = r.options.TemplateInfo
	data["template-path"] = r.options.TemplatePath
	return data
}

// MakeResultEvent creates a result event from internal wrapped event
func (r *Request) MakeResultEvent(wrapped *output2.InternalWrappedEvent) []*output2.ResultEvent {
	if len(wrapped.OperatorsResult.DynamicValues) > 0 && !wrapped.OperatorsResult.Matched {
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
		Type:             "http",
		Host:             types2.ToString(wrapped.InternalEvent["host"]),
		Matched:          types2.ToString(wrapped.InternalEvent["matched"]),
		Metadata:         wrapped.OperatorsResult.PayloadValues,
		ExtractedResults: wrapped.OperatorsResult.OutputExtracts,
		Timestamp:        time.Now(),
		IP:               types2.ToString(wrapped.InternalEvent["ip"]),
		Request:          types2.ToString(wrapped.InternalEvent["request"]),
		Response:         types2.ToString(wrapped.InternalEvent["response"]),
	}
	return data
}
