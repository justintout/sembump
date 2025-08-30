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
	tests := []struct {
		version  string
		expected string
	}{
		{"0.0.0", "0.0.1"},
		{"v1.2.3", "v1.2.4"},
		{"V5.63.632462", "V5.63.632463"},
	}
	for _, tt := range tests {
		bumped, err := Bump(tt.version, BumpOptions{Kind: KindPatch})
		if err != nil {
			t.Errorf("case %q unexpected error: %v", tt.version, err)
		}
		if bumped != tt.expected {
			t.Errorf("got: %s, want: %s", bumped, tt.expected)
		}
	}
}

func testMinor(t *testing.T) {
	tests := []struct {
		version  string
		expected string
	}{
		{"0.0.0", "0.1.0"},
		{"v1.2.3", "v1.3.0"},
		{"V5.63.632462", "V5.64.0"},
	}
	for _, tt := range tests {
		bumped, err := Bump(tt.version, BumpOptions{Kind: KindMinor})
		if err != nil {
			t.Errorf("case %q unexpected error: %v", tt.version, err)
		}
		if bumped != tt.expected {
			t.Errorf("got: %s, want: %s", bumped, tt.expected)
		}
	}
}

func testMajor(t *testing.T) {
	tests := []struct {
		version  string
		expected string
	}{
		{"0.0.0", "1.0.0"},
		{"v1.2.3", "v2.0.0"},
		{"V5.63.632462", "V6.0.0"},
	}
	for _, tt := range tests {
		bumped, err := Bump(tt.version, BumpOptions{Kind: KindMajor})
		if err != nil {
			t.Errorf("case %q unexpected error: %v", tt.version, err)
		}
		if bumped != tt.expected {
			t.Errorf("error bumping patch, got: %s, want: %s", bumped, tt.expected)
		}
	}
}

func testPre(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		opts     BumpOptions
		expected string
	}{
		{"init rc on patch", "1.2.3", BumpOptions{Kind: KindPatch, Prerelease: true}, "1.2.4-rc.1"},
		{"init rc on minor", "1.2.3", BumpOptions{Kind: KindMinor, Prerelease: true}, "1.3.0-rc.1"},
		{"init rc on major", "1.2.3", BumpOptions{Kind: KindMajor, Prerelease: true}, "2.0.0-rc.1"},
		{"increment rc tag.number", "1.2.3-rc.1", BumpOptions{Kind: KindPatch, Prerelease: true}, "1.2.3-rc.2"},
		{"increment number only", "1.2.3-1", BumpOptions{Kind: KindPatch, Prerelease: true}, "1.2.3-2"},
		{"strip pre when not prerelease bump", "1.2.3-rc.4", BumpOptions{Kind: KindPatch, Prerelease: false}, "1.2.4"},
		{"preserve v prefix", "v1.2.3-rc.1", BumpOptions{Kind: KindPatch, Prerelease: true}, "v1.2.3-rc.2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bumped, err := Bump(tt.version, tt.opts)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if bumped != tt.expected {
				t.Fatalf("got %s want %s", bumped, tt.expected)
			}
		})
	}
}

func testInvalid(t *testing.T) {
	tests := []struct {
		name    string
		version string
		opts    BumpOptions
	}{
		{"bad kind", "0.0.0", BumpOptions{Kind: Kind("asdf")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Bump(tt.version, tt.opts)
			if err == nil {
				t.Fatalf("expected error")
			}
		})
	}
}
