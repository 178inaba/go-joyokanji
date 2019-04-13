package joyokanji

import "testing"

func TestList(t *testing.T) {
	rs, err := List()
	if err != nil {
		t.Fatalf("Should not be fail: %v.", err)
	}

	if got, want := len(rs), 2136; got != want {
		t.Fatalf("`List()` length is %d, want %d.", got, want)
	}
}
