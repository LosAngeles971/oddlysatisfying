package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Line struct {
	Line        string
	BeforeDelay int
	AfterDelay  int
}

type Log struct {
	Log         []Line
	Rnd         bool
	BeforeDelay int
	AfterDelay  int
}

type LogOption func(*Log)

func WithRandom() LogOption {
	return func(l *Log) {
		l.Rnd = true
	}
}

func WithBefore(before int) LogOption {
	return func(l *Log) {
		l.BeforeDelay = before
	}
}

func WithAfter(after int) LogOption {
	return func(l *Log) {
		l.AfterDelay = after
	}
}

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

func (l *Log) addf(format string, args ...interface{}) {
	l.Log = append(l.Log, Line{Line: fmt.Sprintf(format, args...), BeforeDelay: 0, AfterDelay: 0})
}

func (l *Log) add(args ...interface{}) {
	l.Log = append(l.Log, Line{Line: fmt.Sprint(args...), BeforeDelay: 0, AfterDelay: 0})
}

func (l *Log) merge(ll *Log) {
	for _, line := range ll.Log {
		l.add(line.Line)
	}
}

func sleep(millis int, rnd bool) {
	var d time.Duration
	if rnd {
		d = time.Duration(rand.Intn(millis)) * time.Millisecond
	} else {
		d = time.Duration(millis) * time.Millisecond
	}
	time.Sleep(d)
}

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