package format

import (
	"io"

	"github.com/lucor/goinfo/internal/report"
)

// Writer is the interface that wraps the Write method
type Writer interface {
	// Write writes the []Report to io.Writer
	Write(io.Writer, []report.Report) error
}
