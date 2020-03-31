package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/lucor/goinfo"
	"github.com/lucor/goinfo/format"
	"github.com/lucor/goinfo/report"
)

var (
	workDir   string
	formatOut string
	version   bool
)

func main() {
	flag.Usage = printHelp
	flag.StringVar(&workDir, "work-dir", "", "")
	flag.StringVar(&formatOut, "format", "text", "")
	flag.BoolVar(&version, "version", false, "")
	flag.Parse()

	if version {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Fprintln(os.Stderr, "Could not retrieve version information. Please ensure module support is activated and build again.")
			os.Exit(1)
		}
		fmt.Println("goinfo version", info.Main.Version)
		os.Exit(0)
	}

	module := flag.Arg(0)

	var f goinfo.Formatter
	switch formatOut {
	case "text":
		f = &format.Text{}
	case "html":
		f = &format.HTMLDetails{}
	case "json":
		f = &format.JSON{}
	default:
		fmt.Fprintln(os.Stderr, "Invalid value for the format flag:", formatOut)
		os.Exit(1)
	}

	workDir, err := ensureWorkDir(workDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	reporters := []goinfo.Reporter{
		&report.GoVersion{},
		&report.GoMod{WorkDir: workDir, Module: module},
		&report.OS{},
		&report.GoEnv{},
	}

	err = goinfo.Write(os.Stdout, reporters, f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
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
  -version          Display the version
  -help             Display this help text
`)
	os.Exit(0)
}
