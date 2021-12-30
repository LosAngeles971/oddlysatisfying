package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Line represents a single entry (or line) of a complete log
type Line struct {
	Line        string // single line of log
	BeforeDelay int    // the delay (in millis) before typing the line
	AfterDelay  int    // the delay (in millis) after typed the line
}

// Log represents an entire log, structured into an array of lines
type Log struct {
	Log         []Line // lines of the log
	Rnd         bool   // false if delays (before and after) mean fixed number of millis, true if delays mean the maximum for a random number
	BeforeDelay int    // default delay (in millis) before typing each line
	AfterDelay  int    // default delay (in millis) after typed each line
}

// Options to characterize the creation of a Log struct
type LogOption func(*Log)

// This option enable the random feature of Log
func WithRandom() LogOption {
	return func(l *Log) {
		l.Rnd = true
	}
}

// This option sets the default before delay of Log
func WithBefore(before int) LogOption {
	return func(l *Log) {
		l.BeforeDelay = before
	}
}

// This option sets the default after delay of Log
func WithAfter(after int) LogOption {
	return func(l *Log) {
		l.AfterDelay = after
	}
}

// New creates a new Log
func New(opts ...LogOption) *Log {
	l := &Log{
		Log:         []Line{},
		Rnd:         false,
		BeforeDelay: 0,
		AfterDelay:  0,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// addf adds one single line to a Log
func (l *Log) addf(format string, args ...interface{}) {
	l.Log = append(l.Log, Line{Line: fmt.Sprintf(format, args...), BeforeDelay: 0, AfterDelay: 0})
}

// add adds one single line to a Log
func (l *Log) add(args ...interface{}) {
	l.Log = append(l.Log, Line{Line: fmt.Sprint(args...), BeforeDelay: 0, AfterDelay: 0})
}

// merge adds lines to Log from the given external Log
func (l *Log) merge(ll *Log) {
	for _, line := range ll.Log {
		l.add(line.Line)
	}
}

// sleep just sleeps for some millis
func sleep(millis int, rnd bool) {
	var d time.Duration
	if rnd {
		d = time.Duration(rand.Intn(millis)) * time.Millisecond
	} else {
		d = time.Duration(millis) * time.Millisecond
	}
	time.Sleep(d)
}

// toScreen types the Log to the standard output
func (l *Log) toScreen() {
	for _, ll := range l.Log {
		if l.BeforeDelay > 0 {
			sleep(l.BeforeDelay, l.Rnd)
		}
		fmt.Println(ll.Line)
		if l.AfterDelay > 0 {
			sleep(l.AfterDelay, l.Rnd)
		}
	}
}
