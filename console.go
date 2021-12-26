package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	mode_console = 0
	mode_logrus = 1
	level_info = 0
	level_error = 1
	level_debug = 2
)

type Console struct {
	delay int
	rnd bool
	mode int
}

type ConsoleOption func(*Console)

func WithDelay(delay int, rnd bool) ConsoleOption {
	return func(c *Console) {
		c.delay = delay
		c.rnd = rnd
		c.mode = mode_console
	}
}

func WithMode(mode int) ConsoleOption {
	return func(c *Console) {
		c.mode = mode
	}
}

func New(opts ...ConsoleOption) Console {
	c := Console{
		delay: 0,
		rnd: false,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

func (c Console) sleep() {
	var d time.Duration
	if c.rnd {
		d = time.Duration(rand.Intn(c.delay)) * time.Millisecond
	} else {
		d = time.Duration(c.delay) * time.Millisecond
	}
	time.Sleep(d)
}

func (c Console) out(l int, args ...interface{}) {
	switch c.mode {
	case mode_logrus:
		switch l {
		case level_debug:
			logrus.Debug(args...)
		case level_error:
			logrus.Error(args...)
		default:
			logrus.Info(args...)
		}
	default:
		fmt.Println(args...)
	}
}

func (c Console) outf(l int, format string, args ...interface{}) {
	switch c.mode {
	case mode_logrus:
		switch l {
		case level_debug:
			logrus.Debugf(format, args...)
		case level_error:
			logrus.Errorf(format, args...)
		default:
			logrus.Infof(format, args...)
		}
	default:
		fmt.Printf(format, args...)
		fmt.Println()
	}
}

func (c Console) Info(args ...interface{}) {
	c.out(level_info, args...)
	c.sleep()
}

func (c Console) Infof(format string, args ...interface{}) {
	c.outf(level_info, format, args...)
	c.sleep()
}

func (c Console) Debug(args ...interface{}) {
	c.out(level_debug, args...)
	c.sleep()
}

func (c Console) Debugf(format string, args ...interface{}) {
	c.outf(level_debug, format, args...)
	c.sleep()
}

func (c Console) Error(args ...interface{}) {
	c.out(level_error, args...)
	c.sleep()
}