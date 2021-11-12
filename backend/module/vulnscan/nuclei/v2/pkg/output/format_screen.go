package output

import (
	types2 "backend/module/vulnscan/nuclei/v2/pkg/types"
	"bytes"
)

// formatScreen formats the output for showing on screen.
func (w *StandardWriter) formatScreen(output *ResultEvent) []byte {
	builder := &bytes.Buffer{}

	if !w.noMetadata {
		//if !w.noTimestamp {
		//	builder.WriteRune('|')
		//	builder.WriteString(output.Timestamp.Format("2006-01-02 15:04:05"))
		//	builder.WriteString("| ")
		//}
		builder.WriteRune('[')
		builder.WriteString(output.TemplateID)

		if output.MatcherName != "" {
			builder.WriteString(":")
			builder.WriteString(output.MatcherName)
		} else if output.ExtractorName != "" {
			builder.WriteString(":")
			builder.WriteString(output.ExtractorName)
		}

		builder.WriteString("]       [")
		builder.WriteString(output.Type)
		builder.WriteString("] ")

		builder.WriteString("        [")
		builder.WriteString(output.Info.SeverityHolder.Severity.String())
		builder.WriteString("] ")
	}
	builder.WriteString(output.Matched)

	// If any extractors, write the results
	if len(output.ExtractedResults) > 0 {
		builder.WriteString("      [")

		for i, item := range output.ExtractedResults {
			builder.WriteString(item)

			if i != len(output.ExtractedResults)-1 {
				builder.WriteRune(',')
			}
		}
		builder.WriteString("]")
	}

	// Write meta if any
	if len(output.Metadata) > 0 {
		builder.WriteString("       [")

		first := true
		for name, value := range output.Metadata {
			if !first {
				builder.WriteRune(',')
			}
			first = false

			builder.WriteString(name)
			builder.WriteRune('=')
			builder.WriteString(types2.ToString(value))
		}
		builder.WriteString("]     ")
	}
	return builder.Bytes()
}
