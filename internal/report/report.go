package report

// Reporter is the interface that wraps the Summary and Info method methods
// along with the ErrorReporter interface
type Reporter interface {
	// Summary returns the summary's report
	Summary() string
	// Info returns the collected info
	Info() (map[string]interface{}, error)
}

// Generate generates the report from a reporter
func Generate(reporter Reporter) Report {
	var e string
	info, err := reporter.Info()
	if err != nil {
		e = err.Error()
	}
	return Report{
		Summary: reporter.Summary(),
		Info:    info,
		Error:   e,
	}
}

// Report reprents a report
type Report struct {
	Summary string                 `json:"summary"`
	Info    map[string]interface{} `json:"info,omitempty"`
	Error   string                 `json:"error,omitempty"`
}
