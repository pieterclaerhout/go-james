package common

import (
	"bytes"
	"text/template"
)

// Template is what can be injected into a subcommand when you need text templates
type Template struct {
	FileSystem
}

// RenderTextTemplate renders a text template
func (t Template) RenderTextTemplate(templateData string, params interface{}) (string, error) {

	tpl := template.New("t")
	parsedTemplate, err := tpl.Parse(templateData)
	if err != nil {
		return "", err
	}

	var renderedTemplate bytes.Buffer
	if err := parsedTemplate.Execute(&renderedTemplate, params); err != nil {
		return "", err
	}
	return renderedTemplate.String(), nil

}

// WriteTextTemplateIfNotExists renders a text template and writes it to a file if it doesn't exist yet
func (t Template) WriteTextTemplateIfNotExists(path string, templateData string, params interface{}) error {

	renderedTemplate, err := t.RenderTextTemplate(templateData, params)
	if err != nil {
		return err
	}

	return t.WriteTextFileIfNotExists(path, renderedTemplate)

}
