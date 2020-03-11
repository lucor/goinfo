package report

import (
	"fmt"
	"os/exec"
	"strings"
)

// GoVersion collects the info about the Go version using the go version command
type GoVersion struct{}

// Summary return the summary
func (i *GoVersion) Summary() string {
	return "Go version info"
}

// Info returns the collected info
func (i *GoVersion) Info() (map[string]interface{}, error) {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not detect go version info: %w", err)
	}

	s := strings.TrimRight(string(out), "\n")
	info := map[string]interface{}{"version": string(s[11:])}
	return info, nil
}
