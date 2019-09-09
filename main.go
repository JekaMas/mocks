package main

import (
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, func() { time.Sleep(time.Second) })
}
