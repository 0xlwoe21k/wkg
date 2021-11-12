package colorizer

import (
	severity2 "backend/module/vulnscan/nuclei/v2/pkg/model/types/severity"
	"fmt"

	"github.com/logrusorgru/aurora"

	"github.com/projectdiscovery/gologger"
)

const (
	fgOrange uint8 = 208
)

func GetColor(colorizer aurora.Aurora, templateSeverity fmt.Stringer) string {
	var method func(arg interface{}) aurora.Value
	switch templateSeverity {
	case severity2.Info:
		method = colorizer.Blue
	case severity2.Low:
		method = colorizer.Green
	case severity2.Medium:
		method = colorizer.Yellow
	case severity2.High:
		method = func(stringValue interface{}) aurora.Value { return colorizer.Index(fgOrange, stringValue) }
	case severity2.Critical:
		method = colorizer.Red
	default:
		gologger.Warning().Msgf("The '%s' severity does not have an color associated!", templateSeverity)
		method = colorizer.White
	}

	return method(templateSeverity.String()).String()
}

func New(colorizer aurora.Aurora) func(severity2.Severity) string {
	return func(severity severity2.Severity) string {
		return GetColor(colorizer, severity)
	}
}
