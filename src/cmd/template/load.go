package template

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

func LoadTemplate(name string) (*template.Template, error) {
	return template.ParseFS(templatesFS, fmt.Sprintf("templates/%s.tmpl", name))
}
