package report

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/lucor/goinfo"
)

// Info returns the collected info
func (i *OS) Info() (goinfo.Info, error) {
	info, err := i.swVers()
	if err != nil {
		return nil, err
	}

	arch, err := i.architecture()
	if err != nil {
		return info, err
	}
	info["architecture"] = arch

	kernel, err := i.kernel()
	if err != nil {
		return info, err
	}
	info["kernel"] = kernel
	return info, err
}

func (i *OS) swVers() (goinfo.Info, error) {
	cmd := exec.Command("sw_vers")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not detect os info using sw_vers command: %w", err)
	}
	return i.parseSwVersCmdOutput(out)
}

func (i *OS) parseSwVersCmdOutput(data []byte) (goinfo.Info, error) {
	// fitlerKeys defines the key to return
	filterKeys := map[string]string{
		"ProductName":    "name",
		"ProductVersion": "version",
		"BuildVersion":   "build_version",
	}
	info := goinfo.Info{}
	buf := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ":")
		if len(tokens) != 2 {
			continue
		}
		key, ok := filterKeys[tokens[0]]
		if !ok {
			continue
		}
		info[key] = strings.Trim(tokens[1], "\t ")
	}
	return info, scanner.Err()
}
