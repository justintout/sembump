package main

import (
	"errors"
	"fmt"
	"strings"

	semver "github.com/blang/semver/v4"
)

// Kind represents the semantic kind of bump.
type Kind string

const (
	KindPatch Kind = "patch"
	KindMinor Kind = "minor"
	KindMajor Kind = "major"
)

var validKinds = []Kind{KindMajor, KindMinor, KindPatch}

// ErrInvalidKind returned when an unsupported kind is requested.
var ErrInvalidKind = errors.New("invalid bump kind")

// BumpOptions customize the bump behavior.
type BumpOptions struct {
	Kind       Kind // required
	Prerelease bool // whether to treat this as a prerelease bump or initialize prerelease
}

// Bump applies a semantic version bump of the provided kind (major/minor/patch) to the
// given version string. It preserves an initial leading 'v' or 'V'.
// When Prerelease is true and the version is already a prerelease of the accepted
// formats ("-number" or "-tag.number"), only the numeric component is incremented.
// When Prerelease is true and the version has no prerelease component, it appends "-rc.1".
func Bump(version string, opts BumpOptions) (string, error) {
	kindLower := strings.ToLower(string(opts.Kind))
	switch Kind(kindLower) {
	case KindPatch, KindMinor, KindMajor:
		// ok
	default:
		return "", fmt.Errorf("%w: %s (valid: %s)", ErrInvalidKind, opts.Kind, kindsString())
	}

	hasPrefixLowerV := strings.HasPrefix(version, "v")
	hasPrefixUpperV := !hasPrefixLowerV && strings.HasPrefix(version, "V")
	if hasPrefixLowerV || hasPrefixUpperV {
		version = version[1:]
	}

	v, err := semver.Make(version)
	if err != nil {
		return "", fmt.Errorf("parse version: %w", err)
	}

	// Handle prerelease only increment cases first.
	if opts.Prerelease && v.Pre != nil {
		// -number
		if len(v.Pre) == 1 && v.Pre[0].IsNum {
			v.Pre[0].VersionNum++
			return render(v, hasPrefixLowerV, hasPrefixUpperV), nil
		}
		// -tag.number
		if len(v.Pre) == 2 && v.Pre[1].IsNum {
			v.Pre[1].VersionNum++
			return render(v, hasPrefixLowerV, hasPrefixUpperV), nil
		}
		return "", fmt.Errorf("unsupported prerelease format: %s", v.Pre)
	}

	// If prerelease flag but no pre component yet, initialize rc.1 after bumping number.
	addPre := opts.Prerelease && v.Pre == nil

	switch Kind(kindLower) {
	case KindPatch:
		v.Patch++
	case KindMinor:
		v.Minor++
		v.Patch = 0
	case KindMajor:
		v.Major++
		v.Minor = 0
		v.Patch = 0
	}

	if addPre {
		s, _ := semver.NewPRVersion("rc")
		n, _ := semver.NewPRVersion("1")
		v.Pre = []semver.PRVersion{s, n}
	} else if !opts.Prerelease && v.Pre != nil {
		// Strip any existing prerelease if not a prerelease bump
		v.Pre = nil
	}

	return render(v, hasPrefixLowerV, hasPrefixUpperV), nil
}

func kindsString() string {
	parts := make([]string, len(validKinds))
	for i, k := range validKinds {
		parts[i] = string(k)
	}
	return strings.Join(parts, " | ")
}

func render(v semver.Version, lower, upper bool) string {
	s := v.String()
	if lower {
		return "v" + s
	}
	if upper {
		return "V" + s
	}
	return s
}
