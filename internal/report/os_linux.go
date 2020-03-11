package report

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Info returns the collected info
func (i *OS) Info() (map[string]interface{}, error) {

	releaseFiles := []string{"/etc/os-release"}
	if matches, err := filepath.Glob("/etc/*release"); err != nil {
		releaseFiles = append(releaseFiles, matches...)
	}
	if matches, err := filepath.Glob("/etc/*version"); err != nil {
		releaseFiles = append(releaseFiles, matches...)
	}

	for _, releaseFile := range releaseFiles {
		b, err := ioutil.ReadFile(releaseFile)
		if err != nil {
			continue
		}
		return i.parseCmdOutput(b)
	}
	return nil, fmt.Errorf("could not found any release file: %#+v", releaseFiles)
}

func (i *OS) parseCmdOutput(data []byte) (map[string]interface{}, error) {
	info := map[string]interface{}{}
	buf := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "=")
		if len(tokens) != 2 {
			continue
		}
		info[tokens[0]] = strings.Trim(tokens[1], `"`)
	}
	return info, scanner.Err()
}
