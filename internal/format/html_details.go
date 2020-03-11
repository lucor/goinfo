package format

import (
	"fmt"
	"html/template"
	"io"

	"github.com/lucor/goinfo/internal/report"
)

// HTMLDetails writes reports in HTML Details format
type HTMLDetails struct{}

// Write writes the []Reporter to io.Writer
func (w *HTMLDetails) Write(out io.Writer, reports []report.Report) error {
	t := template.Must(template.New("html_details").Parse(htmlDetailsTpl))
	// Execute the template for each report
	for _, r := range reports {
		err := t.Execute(out, r)
		if err != nil {
			return fmt.Errorf("could not execute the html details template: %w", err)
		}
	}
	return nil
}

const htmlDetailsTpl = `
<details><summary>{{.Summary}}</summary><br><pre>
{{ range $key, $value := .Info -}}
{{ $key }}={{ $value }}
{{ end -}}
</pre></details>
`
