package main

import (
	_ "embed"
	"flag"
	"fmt"
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

	if len(flag.Args()) < 1 {
		usageAndExit(1, "must pass a semver string, ex: %s", myVersion)
	}

	kind = strings.ToLower(kind)
	version := flag.Arg(0)
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
