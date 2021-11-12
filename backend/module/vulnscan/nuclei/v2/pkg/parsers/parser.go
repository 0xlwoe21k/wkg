package parsers

import (
	filter2 "backend/module/vulnscan/nuclei/v2/pkg/catalog/loader/filter"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	templates2 "backend/module/vulnscan/nuclei/v2/pkg/templates"
	cache2 "backend/module/vulnscan/nuclei/v2/pkg/templates/cache"
	utils2 "backend/module/vulnscan/nuclei/v2/pkg/utils"
	stats2 "backend/module/vulnscan/nuclei/v2/pkg/utils/stats"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"

	"github.com/projectdiscovery/gologger"
)

const mandatoryFieldMissingTemplate = "mandatory '%s' field is missing"

// LoadTemplate returns true if the template is valid and matches the filtering criteria.
func LoadTemplate(templatePath string, tagFilter *filter2.TagFilter, extraTags []string) (bool, error) {
	template, templateParseError := ParseTemplate(templatePath)
	if templateParseError != nil {
		return false, templateParseError
	}

	if len(template.Workflows) > 0 {
		return false, nil
	}

	templateInfo := template.Info
	if validationError := validateMandatoryInfoFields(&templateInfo); validationError != nil {
		return false, validationError
	}

	return isTemplateInfoMetadataMatch(tagFilter, &templateInfo, extraTags)
}

// LoadWorkflow returns true if the workflow is valid and matches the filtering criteria.
func LoadWorkflow(templatePath string) (bool, error) {
	template, templateParseError := ParseTemplate(templatePath)
	if templateParseError != nil {
		return false, templateParseError
	}

	templateInfo := template.Info

	if len(template.Workflows) > 0 {
		if validationError := validateMandatoryInfoFields(&templateInfo); validationError != nil {
			return false, validationError
		}
		return true, nil
	}

	return false, nil
}

func isTemplateInfoMetadataMatch(tagFilter *filter2.TagFilter, templateInfo *model2.Info, extraTags []string) (bool, error) {
	templateTags := templateInfo.Tags.ToSlice()
	templateAuthors := templateInfo.Authors.ToSlice()
	templateSeverity := templateInfo.SeverityHolder.Severity

	match, err := tagFilter.Match(templateTags, templateAuthors, templateSeverity, extraTags)

	if err == filter2.ErrExcluded {
		return false, filter2.ErrExcluded
	}

	return match, err
}

func validateMandatoryInfoFields(info *model2.Info) error {
	if info == nil {
		return fmt.Errorf(mandatoryFieldMissingTemplate, "info")
	}

	if utils2.IsBlank(info.Name) {
		return fmt.Errorf(mandatoryFieldMissingTemplate, "name")
	}

	if info.Authors.IsEmpty() {
		return fmt.Errorf(mandatoryFieldMissingTemplate, "author")
	}
	return nil
}

var (
	parsedTemplatesCache *cache2.Templates
	ShouldValidate       bool
	fieldErrorRegexp     = regexp.MustCompile(`not found in`)
)

const (
	SyntaxWarningStats = "syntax-warnings"
	SyntaxErrorStats   = "syntax-errors"
)

func init() {

	parsedTemplatesCache = cache2.New()

	stats2.NewEntry(SyntaxWarningStats, "Found %d templates with syntax warning (use -validate flag for further examination)")
	stats2.NewEntry(SyntaxErrorStats, "Found %d templates with syntax error (use -validate flag for further examination)")
}

// ParseTemplate parses a template and returns a *templates.Template structure
func ParseTemplate(templatePath string) (*templates2.Template, error) {
	if value, err := parsedTemplatesCache.Has(templatePath); value != nil {
		return value.(*templates2.Template), err
	}

	f, err := os.Open(templatePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	template := &templates2.Template{}
	if err := yaml.UnmarshalStrict(data, template); err != nil {
		errString := err.Error()
		if !fieldErrorRegexp.MatchString(errString) {
			stats2.Increment(SyntaxErrorStats)
			return nil, err
		}
		stats2.Increment(SyntaxWarningStats)
		if ShouldValidate {
			gologger.Error().Msgf("Syntax warnings for template %s: %s", templatePath, err)
		} else {
			gologger.Warning().Msgf("Syntax warnings for template %s: %s", templatePath, err)
		}
	}
	parsedTemplatesCache.Store(templatePath, template, nil)
	return template, nil
}
