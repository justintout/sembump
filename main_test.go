package main

import "testing"

func TestBump(t *testing.T) {
	t.Run("Patch", testPatch)
	t.Run("Minor", testMinor)
	t.Run("Major", testMajor)
	t.Run("Pre", testPre)
	t.Run("Invalid", testInvalid)
}

func testPatch(t *testing.T) {
	kind := "patch"
	tests := []struct {
		version  string
		expected string
	}{
		{"0.0.0", "0.0.1"},
		{"v1.2.3", "v1.2.4"},
		{"V5.63.632462", "V5.63.632463"},
	}
	for _, tt := range tests {
		bumped, err := bump(tt.version, kind)
		if err != nil {
			t.Errorf("case %q unexpected error: %v", tt.version, err)
		}
		if bumped != tt.expected {
			t.Errorf("got: %s, want: %s", bumped, tt.expected)
		}
	}
}

func testMinor(t *testing.T) {
	kind := "minor"
	tests := []struct {
		version  string
		expected string
	}{
		{"0.0.0", "0.1.0"},
		{"v1.2.3", "v1.3.0"},
		{"V5.63.632462", "V5.64.0"},
	}
	for _, tt := range tests {
		bumped, err := bump(tt.version, kind)
		if err != nil {
			t.Errorf("case %q unexpected error: %v", tt.version, err)
		}
		if bumped != tt.expected {
			t.Errorf("got: %s, want: %s", bumped, tt.expected)
		}
	}
}

func testMajor(t *testing.T) {
	kind := "major"
	tests := []struct {
		version  string
		expected string
	}{
		{"0.0.0", "1.0.0"},
		{"v1.2.3", "v2.0.0"},
		{"V5.63.632462", "V6.0.0"},
	}
	for _, tt := range tests {
		bumped, err := bump(tt.version, kind)
		if err != nil {
			t.Errorf("case %q unexpected error: %v", tt.version, err)
		}
		if bumped != tt.expected {
			t.Errorf("error bumping patch, got: %s, want: %s", bumped, tt.expected)
		}
	}
}

func testPre(t *testing.T) {
	t.Skip("prerelease bump tests not written")
}

func testInvalid(t *testing.T) {
	tests := []struct {
		kind    string
		version string
	}{
		{"patch", "0.0.0"},
		{"minor", "v1.2.3"},
		{"major", "V5.63.632462"},
		{"asdf", "0.0.0"},
	}
	for _, tt := range tests {
		_, err := bump(tt.version, kind)
		if err == nil {
			t.Errorf("expected error")
		}
	}
}
