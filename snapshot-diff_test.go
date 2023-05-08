package main

import (
	"testing"
)

func Test(t *testing.T) {
	got := 4
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
