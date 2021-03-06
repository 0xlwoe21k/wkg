package main

import (
	templates2 "backend/module/vulnscan/nuclei/v2/pkg/templates"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/alecthomas/jsonschema"
)

var pathRegex = regexp.MustCompile(`backend/services/vulnscan/nuclei/v2/(?:internal|pkg)/(?:.*/)?([A-Za-z\.]+)`)

func main() {
	// Generate yaml syntax documentation
	data, err := templates2.GetTemplateDoc().Encode()
	if err != nil {
		log.Fatalf("Could not encode docs: %s\n", err)
	}
	err = ioutil.WriteFile(os.Args[1], data, 0777)
	if err != nil {
		log.Fatalf("Could not write docs: %s\n", err)
	}

	// Generate jsonschema
	r := &jsonschema.Reflector{
		PreferYAMLSchema:      true,
		YAMLEmbeddedStructs:   true,
		FullyQualifyTypeNames: true,
	}
	jsonschemaData := r.Reflect(&templates2.Template{})

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(jsonschemaData)

	schema := buf.String()
	for _, match := range pathRegex.FindAllStringSubmatch(schema, -1) {
		schema = strings.ReplaceAll(schema, match[0], match[1])
	}
	err = ioutil.WriteFile(os.Args[2], []byte(schema), 0777)
	if err != nil {
		log.Fatalf("Could not write jsonschema: %s\n", err)
	}
}
