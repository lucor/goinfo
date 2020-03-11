package report

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Info returns the collected info
func (i *OS) Info() (map[string]interface{}, error) {
	cmd := exec.Command("sw_vers")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not detect os info using sw_vers command: %w", err)
	}
	return i.parseCmdOutput(out)
}

func (i *OS) parseCmdOutput(data []byte) (map[string]interface{}, error) {
	info := map[string]interface{}{}
	buf := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ":")
		if len(tokens) != 2 {
			continue
		}
		info[tokens[0]] = strings.Trim(tokens[1], "\t ")
	}
	return info, scanner.Err()
}
