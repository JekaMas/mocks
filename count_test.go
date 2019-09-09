package main

import (
	"bytes"
	"testing"
)

type SpySleeper func()

func newSpySleeper() (SpySleeper, *int) {
	c := new(int)
	return func() { (*c)++ }, c
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	spySleeper, calls := newSpySleeper()

	Countdown(buffer, spySleeper)

	if *calls != 4 {
		t.Errorf("got time %d want %d", *calls, 4)
	}

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
