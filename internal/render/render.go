package render

import (
	"bytes"
	_ "embed"

	"text/template"
)

//go:embed templates/options.go.tmpl
var optionsTemplate []byte

func Render(f *File) (*bytes.Buffer, error) {
	tmpl, err := template.New("options").Parse(string(optionsTemplate))
	if err != nil {
		return nil, err
	}
	b := &bytes.Buffer{}
	if err := tmpl.Execute(b, f); err != nil {
		return nil, err
	}
	return b, nil
}
