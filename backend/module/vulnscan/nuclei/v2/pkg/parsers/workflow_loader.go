package parsers

import (
	filter2 "backend/module/vulnscan/nuclei/v2/pkg/catalog/loader/filter"
	model2 "backend/module/vulnscan/nuclei/v2/pkg/model"
	protocols2 "backend/module/vulnscan/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/gologger"
)

type workflowLoader struct {
	pathFilter *filter2.PathFilter
	tagFilter  *filter2.TagFilter
	options    *protocols2.ExecuterOptions
}

// NewLoader returns a new workflow loader structure
func NewLoader(options *protocols2.ExecuterOptions) (model2.WorkflowLoader, error) {
	tagFilter := filter2.New(&filter2.Config{
		Tags:        options.Options.Tags,
		ExcludeTags: options.Options.ExcludeTags,
		Authors:     options.Options.Author,
		Severities:  options.Options.Severities,
		IncludeTags: options.Options.IncludeTags,
	})
	pathFilter := filter2.NewPathFilter(&filter2.PathFilterConfig{
		IncludedTemplates: options.Options.IncludeTemplates,
		ExcludedTemplates: options.Options.ExcludedTemplates,
	}, options.Catalog)
	return &workflowLoader{pathFilter: pathFilter, tagFilter: tagFilter, options: options}, nil
}

func (w *workflowLoader) GetTemplatePathsByTags(templateTags []string) []string {
	includedTemplates := w.options.Catalog.GetTemplatesPath([]string{w.options.Options.TemplatesDirectory})
	templatePathMap := w.pathFilter.Match(includedTemplates)

	loadedTemplates := make([]string, 0, len(templatePathMap))
	for templatePath := range templatePathMap {
		loaded, err := LoadTemplate(templatePath, w.tagFilter, templateTags)
		if err != nil {
			gologger.Warning().Msgf("Could not load template %s: %s\n", templatePath, err)
		} else if loaded {
			loadedTemplates = append(loadedTemplates, templatePath)
		}
	}
	return loadedTemplates
}

func (w *workflowLoader) GetTemplatePaths(templatesList []string, noValidate bool) []string {
	includedTemplates := w.options.Catalog.GetTemplatesPath(templatesList)
	templatesPathMap := w.pathFilter.Match(includedTemplates)

	loadedTemplates := make([]string, 0, len(templatesPathMap))
	for templatePath := range templatesPathMap {
		matched, err := LoadTemplate(templatePath, w.tagFilter, nil)
		if err != nil {
			gologger.Warning().Msgf("Could not load template %s: %s\n", templatePath, err)
		} else if matched || noValidate {
			loadedTemplates = append(loadedTemplates, templatePath)
		}
	}
	return loadedTemplates
}
