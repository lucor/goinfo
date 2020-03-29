package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucor/goinfo/internal/format"
	"github.com/lucor/goinfo/internal/report"
)

var (
	workDir    string
	modulePath string
	formatOut  string
)

func main() {
	flag.Usage = printHelp
	flag.StringVar(&workDir, "work-dir", "", "")
	flag.StringVar(&formatOut, "format", "text", "")
	flag.Parse()

	module := flag.Arg(0)

	var w format.Writer
	switch formatOut {
	case "text":
		w = &format.Text{}
	case "html":
		w = &format.HTMLDetails{}
	case "json":
		w = &format.JSON{}
	default:
		fmt.Println("Invalid value for the format flag:", formatOut)
		os.Exit(1)
	}

	workDir, err := ensureWorkDir(workDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reporters := []report.Reporter{
		&report.GoVersion{},
		&report.GoMod{WorkDir: workDir, Module: module},
		&report.OS{},
		&report.GoEnv{},
	}

	reports := []report.Report{}
	for _, reporter := range reporters {
		reports = append(reports, report.Generate(reporter))
	}

	w.Write(os.Stdout, reports)

	// display errors, if any
	for _, report := range reports {
		err := report.Error
		if err != "" {
			fmt.Fprintf(os.Stderr, "[WARN] Report %q: %s\n", report.Summary, err)
		}
	}
}

func ensureWorkDir(workDir string) (string, error) {
	if workDir == "" {
		workDir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("could not get the path for the current working dir: %w", err)
		}
		return workDir, nil
	}
	workDir, err := filepath.Abs(workDir)
	if err != nil {
		return "", fmt.Errorf("could not get the path for the current working dir: %w", err)
	}
	return workDir, nil
}

func printHelp() {
	fmt.Print(`goinfo:

Usage: goinfo [options...] [module]
  List information about a Go module and the development environment.
  Default for the module in current directory.
Options:
  -work-dir         Path of the working dir. Default to current dir
  -format           Format output for the report. Supported: text, html, json. Default to text
  -help             Display this help text
`)
	os.Exit(0)
}
