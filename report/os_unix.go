// +build darwin linux openbsd freebsd netbsd dragonflypackage report
package report

import (
	"fmt"
	"os/exec"
	"strings"
)

func (i *OS) architecture() (string, error) {
	cmd := exec.Command("uname", "-m")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not detect architecture using uname command: %w", err)
	}

	arch := strings.Trim(string(out), "\n")

	return arch, nil
}

func (i *OS) kernel() (string, error) {
	cmd := exec.Command("uname", "-rsv")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not detect the kernel using uname command: %w", err)
	}

	kernel := strings.Trim(string(out), "\n")

	return kernel, nil
}
