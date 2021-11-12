package expressions

import (
	dsl2 "backend/module/vulnscan/nuclei/v2/pkg/operators/common/dsl"
	generators2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/common/generators"
	replacer2 "backend/module/vulnscan/nuclei/v2/pkg/protocols/common/replacer"
	"regexp"

	"github.com/Knetic/govaluate"
)

var templateExpressionRegex = regexp.MustCompile(`(?m)\{\{[^}]+\}\}["'\)\}]*`)

// Evaluate checks if the match contains a dynamic variable, for each
// found one we will check if it's an expression and can
// be compiled, it will be evaluated and the results will be returned.
//
// The provided keys from finalValues will be used as variable names
// for substitution inside the expression.
func Evaluate(data string, base map[string]interface{}) (string, error) {
	data = replacer2.Replace(data, base)

	dynamicValues := make(map[string]interface{})
	for _, match := range templateExpressionRegex.FindAllString(data, -1) {
		expr := generators2.TrimDelimiters(match)

		compiled, err := govaluate.NewEvaluableExpressionWithFunctions(expr, dsl2.HelperFunctions())
		if err != nil {
			continue
		}
		result, err := compiled.Evaluate(base)
		if err != nil {
			continue
		}
		dynamicValues[expr] = result
	}
	// Replacer dynamic values if any in raw request and parse  it
	return replacer2.Replace(data, dynamicValues), nil
}

// EvaluateByte checks if the match contains a dynamic variable, for each
// found one we will check if it's an expression and can
// be compiled, it will be evaluated and the results will be returned.
//
// The provided keys from finalValues will be used as variable names
// for substitution inside the expression.
func EvaluateByte(data []byte, base map[string]interface{}) ([]byte, error) {
	final := replacer2.Replace(string(data), base)

	dynamicValues := make(map[string]interface{})
	for _, match := range templateExpressionRegex.FindAllString(final, -1) {
		expr := generators2.TrimDelimiters(match)

		compiled, err := govaluate.NewEvaluableExpressionWithFunctions(expr, dsl2.HelperFunctions())
		if err != nil {
			continue
		}
		result, err := compiled.Evaluate(base)
		if err != nil {
			continue
		}
		dynamicValues[expr] = result
	}
	// Replacer dynamic values if any in raw request and parse  it
	return []byte(replacer2.Replace(final, dynamicValues)), nil
}
