package format

import (
	"fmt"
	"io"
	"text/template"

	"github.com/lucor/goinfo/internal/report"
)

// Text writes reports in text format
type Text struct{}

// Write writes the []Reporter to io.Writer
func (w *Text) Write(out io.Writer, reports []report.Report) error {
	t := template.Must(template.New("text").Parse(textTpl))
	// Execute the template for each report
	for _, r := range reports {
		err := t.Execute(out, r)
		if err != nil {
			return fmt.Errorf("could not execute the text template: %w", err)
		}
	}
	return nil
}

const textTpl = `## {{.Summary}}
{{ range $key, $value := .Info -}}
{{ $key }}="{{ $value }}"
{{ end }}
`
