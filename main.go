package main

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

type uselessFunc func()

var vtp = map[int]uselessFunc{
	0: download,
	1: mvs,
}

var clo = map[int]uselessFunc{
	0: standard,
}

func main() {
	rand.Seed(time.Now().UnixNano())
	calls := rand.Intn(20)
	log.Infof("batch size is %v", calls)
	for i := 0; i < calls; i++ {
		call := vtp[rand.Intn(len(vtp))]
		call()
	}
	call := clo[rand.Intn(len(clo))]
	call()
}
