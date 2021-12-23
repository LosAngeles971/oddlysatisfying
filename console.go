package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

type Console struct {
	delay int
	rnddelay int
}

type ConsoleOption func(*Console)

func WithDelay(delay int) ConsoleOption {
	return func(c *Console) {
		c.delay = delay
		c.rnddelay = 0
	}
}

func WithRndDelay(rnddelay int) ConsoleOption {
	return func(c *Console) {
		c.delay = 0
		c.rnddelay = rnddelay
	}
}

func New(opts ...ConsoleOption) Console {
	c := Console{
		delay: 0,
		rnddelay: 0,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

func (c Console) sleep() {
	if c.delay > 0 {
		time.Sleep(time.Duration(c.delay) * time.Second)
	}
	if c.rnddelay > 0 {
		time.Sleep(time.Duration(rand.Intn(c.rnddelay)) * time.Second)
	}
}

func (c Console) Info(args ...interface{}) {
	c.sleep()
	log.Info(args)
}

func (c Console) Infof(args ...interface{}) {
	c.sleep()
	log.Info(args)
}

func (c Console) Debug(args ...interface{}) {
	c.sleep()
	log.Info(args)
}

func (c Console) Error(args ...interface{}) {
	c.sleep()
	log.Info(args)
}