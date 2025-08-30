package main

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// buildTestBinary builds the current module's binary and returns its path.
func buildTestBinary(t *testing.T) string {
	t.Helper()
	tmp := t.TempDir()
	bin := filepath.Join(tmp, "sembump")
	if runtime.GOOS == "windows" { // ensure .exe for windows if tests ever run there
		bin += ".exe"
	}
	cmd := exec.Command("go", "build", "-o", bin, ".")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to build binary: %v\n%s", err, stderr.String())
	}
	return bin
}

func runBinary(t *testing.T, bin string, stdin string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			exitCode = ee.ExitCode()
		} else {
			// Unexpected failure to even start process
			t.Fatalf("command failed to start: %v", err)
		}
	}
	return outBuf.String(), errBuf.String(), exitCode
}

func TestMainStdinImplicit(t *testing.T) {
	bin := buildTestBinary(t)
	stdout, stderr, code := runBinary(t, bin, "1.2.3\n", "-kind", "patch")
	if code != 0 {
		t.Fatalf("expected exit 0 got %d stderr=%s", code, stderr)
	}
	if stdout != "1.2.4" { // patch bump
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestMainStdinExplicitDash(t *testing.T) {
	bin := buildTestBinary(t)
	stdout, stderr, code := runBinary(t, bin, "1.2.3\n", "-kind", "minor", "-")
	if code != 0 {
		t.Fatalf("expected exit 0 got %d stderr=%s", code, stderr)
	}
	if stdout != "1.3.0" { // minor bump
		t.Fatalf("unexpected stdout: %q", stdout)
	}
}

func TestMainStdinExplicitDashEmpty(t *testing.T) {
	bin := buildTestBinary(t)
	stdout, stderr, code := runBinary(t, bin, "", "-")
	if code == 0 {
		t.Fatalf("expected non-zero exit code with empty stdin and '-' arg")
	}
	if stdout != "" {
		t.Fatalf("expected no stdout, got %q", stdout)
	}
	if !strings.Contains(stderr, "stdin empty while '-' specified") {
		t.Fatalf("stderr did not contain expected message; got: %s", stderr)
	}
}
