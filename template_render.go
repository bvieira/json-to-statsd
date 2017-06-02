package main

import (
	"bytes"
	"errors"
	"html/template"
	"os"
	"strings"
)

//TemplateRender apply template rules
type TemplateRender struct {
	funcs template.FuncMap
}

//NewTemplateRender TemplateRender constructor
func NewTemplateRender() *TemplateRender {
	funcMap := template.FuncMap{
		"env":   getEnv,
		"envd":  getEnvDefault,
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
	}

	return &TemplateRender{funcs: funcMap}
}

//Render renders a template
func (t *TemplateRender) Render(tpl string) (string, error) {
	parsedTemplate, err := template.New("config").Funcs(t.funcs).Parse(tpl)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = parsedTemplate.Execute(&result, nil)
	return result.String(), err
}

func getEnv(key string) (string, error) {
	result, exists := os.LookupEnv(key)
	if !exists {
		return "", errors.New("env not found")
	}
	return result, nil
}

func getEnvDefault(key, defaultValue string) string {
	result, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return result
}
