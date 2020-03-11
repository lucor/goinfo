package format

import (
	"encoding/json"
	"io"

	"github.com/lucor/goinfo/internal/report"
)

// JSON writes reports in JSON format
type JSON struct{}

// Write writes the []Reporter to io.Writer
func (w *JSON) Write(out io.Writer, reports []report.Report) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "\t")
	return enc.Encode(reports)
}
