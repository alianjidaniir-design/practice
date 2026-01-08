package main

import (
	"testing"
	"unicode/utf8"
)

func TestR1(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{" ", " "},
		{"!12345@", "@54321!"},
		{"Mastering Go", "oG gniretsaM"},
	}
	for _, tc := range testCases {
		rev, err := R1(tc.in)
		if err != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("R1 = %q, want %q", rev, tc.want)
		}
	}
}
func TestR2(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{" ", " "},
		{"!12345@", "@54321!"},
		{"Mastering Go", "oG gniretsaM"},
	}
	for _, tc := range testCases {
		rev, err := R2(tc.in)
		if err != nil {
			return
		}
		if rev != tc.want {
			t.Errorf("R2 = %q, want %q", rev, tc.want)
		}
	}
}

func FuzzR1(f *testing.F) {
	testcases := []string{"Hello, World", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := R1(orig)
		if err != nil {
			return
		}
		doubleRev, err := R1(rev)
		if err != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("R1 = %q, want %q", rev, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(string(rev)) {
			t.Errorf("utf8.ValidString = %v, want valid utf8 string", rev)
		}
	})

}

func FuzzR2(f *testing.F) {
	testCases := []string{"Hello, World", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := R2(orig)
		if err != nil {
			return
		}
		doubleRev, err := R2(rev)
		if err != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("R2 = %q, want %q", rev, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("utf8.ValidString = %v, want valid utf8 string", rev)
		}
	})
}
