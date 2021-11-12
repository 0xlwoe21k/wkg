package file

import (
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	extractors2 "backend/module/vulnscan/nuclei/v2/pkg/operators/extractors"
	matchers2 "backend/module/vulnscan/nuclei/v2/pkg/operators/matchers"
	output2 "backend/module/vulnscan/nuclei/v2/pkg/output"
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"bufio"
	"strings"
	"time"
)

// Match matches a generic data response again a given matcher
func (r *Request) Match(data map[string]interface{}, matcher *matchers2.Matcher) bool {
	partString := matcher.Part
	switch partString {
	case "body", "all", "data", "":
		partString = "raw"
	}

	item, ok := data[partString]
	if !ok {
		return false
	}
	itemStr := types2.ToString(item)

	switch matcher.GetType() {
	case matchers2.SizeMatcher:
		return matcher.Result(matcher.MatchSize(len(itemStr)))
	case matchers2.WordsMatcher:
		return matcher.Result(matcher.MatchWords(itemStr))
	case matchers2.RegexMatcher:
		return matcher.Result(matcher.MatchRegex(itemStr))
	case matchers2.BinaryMatcher:
		return matcher.Result(matcher.MatchBinary(itemStr))
	case matchers2.DSLMatcher:
		return matcher.Result(matcher.MatchDSL(data))
	}
	return false
}

// Extract performs extracting operation for an extractor on model and returns true or false.
func (r *Request) Extract(data map[string]interface{}, extractor *extractors2.Extractor) map[string]struct{} {
	partString := extractor.Part
	switch partString {
	case "body", "all", "data", "":
		partString = "raw"
	}

	item, ok := data[partString]
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
func (r *Request) responseToDSLMap(raw, host, matched string) output2.InternalEvent {
	data := make(output2.InternalEvent, 5)

	// Some data regarding the request metadata
	data["path"] = host
	data["matched"] = matched
	data["raw"] = raw
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
	raw, ok := wrapped.InternalEvent["raw"]
	if !ok {
		return results
	}
	rawStr, ok := raw.(string)
	if !ok {
		return results
	}

	// Identify the position of match in file using a dirty hack.
	for _, result := range results {
		for _, extraction := range result.ExtractedResults {
			scanner := bufio.NewScanner(strings.NewReader(rawStr))

			line := 1
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), extraction) {
					if result.FileToIndexPosition == nil {
						result.FileToIndexPosition = make(map[string]int)
					}
					result.FileToIndexPosition[result.Matched] = line
					continue
				}
				line++
			}
		}
	}
	return results
}

func (r *Request) makeResultEventItem(wrapped *output2.InternalWrappedEvent) *output2.ResultEvent {
	data := &output2.ResultEvent{
		TemplateID:       types2.ToString(wrapped.InternalEvent["template-id"]),
		TemplatePath:     types2.ToString(wrapped.InternalEvent["template-path"]),
		Info:             wrapped.InternalEvent["template-info"].(model2.Info),
		Type:             "file",
		Path:             types2.ToString(wrapped.InternalEvent["path"]),
		Matched:          types2.ToString(wrapped.InternalEvent["matched"]),
		Host:             types2.ToString(wrapped.InternalEvent["matched"]),
		ExtractedResults: wrapped.OperatorsResult.OutputExtracts,
		Response:         types2.ToString(wrapped.InternalEvent["raw"]),
		Timestamp:        time.Now(),
	}
	return data
}
