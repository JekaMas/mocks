package main

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleep func()) {
	for i := countdownStart; i > 0; i-- {
		sleep()
		_, _ = fmt.Fprintln(out, i)
	}

	sleep()
	_, _ = fmt.Fprint(out, finalWord)
}

