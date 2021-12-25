package main

import (
	"math/rand"
	"time"
)

type uselessFunc func()

var vtp = map[int]uselessFunc{
	0: download,
	1: mvs,
	2: bootVax,
	3: c64,
	4: adminports,
}

var clo = map[int]uselessFunc{
	0: standard,
}

func main() {
	rand.Seed(time.Now().UnixNano())
	calls := rand.Intn(5)
	for i := 0; i < calls; i++ {
		call := vtp[rand.Intn(len(vtp))]
		call()
	}
	call := clo[rand.Intn(len(clo))]
	call()
}
