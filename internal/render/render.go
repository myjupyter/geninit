package render

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed templates/options.go.tmpl
var optionsTemplate []byte

func Render(data any) (*bytes.Buffer, error) {
	tmpl, err := template.New("options").Parse(string(optionsTemplate))
	if err != nil {
		return nil, err
	}
	b := &bytes.Buffer{}
	if err := tmpl.Execute(b, data); err != nil {
		return nil, err
	}
	return b, nil
}
