package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	myName = "sembump"

	defaultKind = "patch"
)

var (
	//go:embed VERSION
	myVersion string

	kind  string
	pre   bool
	kinds = []string{"major", "minor", "patch"}
)

func main() {
	flag.StringVar(&kind, "kind", defaultKind, fmt.Sprintf("Kind of version bump [%s]", strings.Join(kinds, " | ")))
	flag.BoolVar(&pre, "pre", false, "Bump as prerelease version")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s %s\n", myName, myVersion)
		flag.PrintDefaults()
	}
	flag.Parse()

	var version string
	if len(flag.Args()) >= 1 { // explicit argument provided
		arg := flag.Arg(0)
		if arg == "-" { // force read from stdin
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				die("failed to read stdin: %v", err)
			}
			version = strings.TrimSpace(string(data))
			if version == "" {
				usageAndExit(1, "stdin empty while '-' specified; pipe a version, ex: echo %s | %s -kind patch -", myVersion, myName)
			}
		} else {
			version = arg
		}
	} else {
		// No positional argument: attempt to read from stdin only if it's a pipe/file.
		if fi, err := os.Stdin.Stat(); err == nil && (fi.Mode()&os.ModeCharDevice) == 0 {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				die("failed to read stdin: %v", err)
			}
			version = strings.TrimSpace(string(data))
		}
		if version == "" { // still empty -> error
			usageAndExit(1, "must provide a semver via arg, '-', or piped on stdin, ex: echo %s | %s -kind patch", myVersion, myName)
		}
	}

	kind = strings.ToLower(kind)
	bumped, err := Bump(version, BumpOptions{Kind: Kind(kind), Prerelease: pre})
	if err != nil {
		die("failed to bump %s: %v", version, err)
	}
	fmt.Fprint(os.Stdout, bumped)
}

func die(message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, message, args...)
	os.Exit(1)
}

func usageAndExit(exitCode int, message string, args ...interface{}) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message, args...)
		fmt.Fprint(os.Stderr, "\n\n")
	}
	flag.Usage()
	os.Exit(exitCode)
}
