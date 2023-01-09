package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/blang/semver"
)

const (
	myName = "sembump"

	defaultKind = "patch"
)

var (
	//go:embed VERSION
	myVersion string

	kind  string
	kinds = []string{"major", "minor", "patch"}
	pre   bool
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
	bumped, err := bump(version, kind)
	if err != nil {
		die("failed to bump %s: %v", version, err)
	}
	fmt.Fprint(os.Stdout, bumped)
}

func bump(version, kind string) (string, error) {
	var (
		hasPrefixLowerV bool
		hasPrefixUpperV bool

		bumped string
	)

	if hasPrefixLowerV = strings.HasPrefix(version, "v"); hasPrefixLowerV {
		version = strings.TrimPrefix(version, "v")
	}
	if hasPrefixUpperV = strings.HasPrefix(version, "V"); hasPrefixUpperV {
		version = strings.TrimPrefix(version, "V")
	}

	v, err := semver.Make(version)
	if err != nil {
		return bumped, fmt.Errorf("failed to parse version: %v", err)
	}

	switch {
	case !pre && v.Pre != nil:
		v.Pre = nil
	case pre && v.Pre != nil:
		// -number
		if len(v.Pre) == 1 && v.Pre[0].IsNum {
			v.Pre[0].VersionNum++
			break
		}
		// -tag.number
		if len(v.Pre) == 2 && v.Pre[1].IsNum {
			v.Pre[1].VersionNum++
			break
		}
		return bumped, fmt.Errorf(`can't handle prerelease tags not of the form "-tag.number" or "-number"`)
	case kind == "patch":
		if pre {
			s, _ := semver.NewPRVersion("rc")
			n, _ := semver.NewPRVersion("1")
			v.Pre = []semver.PRVersion{s, n}
		}
		v.Patch++
	case kind == "minor":
		if pre {
			s, _ := semver.NewPRVersion("rc")
			n, _ := semver.NewPRVersion("1")
			v.Pre = []semver.PRVersion{s, n}
		}
		v.Minor++
		v.Patch = 0
	case kind == "major":
		if pre {
			s, _ := semver.NewPRVersion("rc")
			n, _ := semver.NewPRVersion("1")
			v.Pre = []semver.PRVersion{s, n}
		}
		v.Major++
		v.Minor = 0
		v.Patch = 0
	default:
		return bumped, fmt.Errorf("%s is not valid, please use one of the following [%s]", kind, strings.Join(kinds, " | "))
	}

	bumped = v.String()

	switch {
	case hasPrefixLowerV:
		return "v" + bumped, nil
	case hasPrefixUpperV:
		return "V" + bumped, nil
	default:
		return bumped, nil
	}
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
