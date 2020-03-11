package report

import (
	"fmt"
	"os/exec"
	"strings"
)

// Info returns the collected info about the OS
func (i *OS) Info() (map[string]interface{}, error) {
	cmd := exec.Command("cmd", "/C", "ver")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not detect os info using ver command: %w", err)
	}

	s := strings.Trim(string(out), "\r\n")

	info := map[string]interface{}{"version": s}
	return info, nil
}
