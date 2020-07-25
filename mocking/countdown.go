package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper that sleeps
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper that can configure sleep time
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep that will sleep for 1 second
func (sleeper *ConfigurableSleeper) Sleep() {
	sleeper.sleep(sleeper.duration)
}

// Countdown countdown from 3 to 1 and prints go
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
