package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type uselessFunc func()

var vtp = map[int]uselessFunc{
	0: download,
	1: mvs,
	2: bootVax,
	3: c64,
	4: adminports,
	5: fsck,
	6: bootMsDos,
	7: stateLinux,
	8: qradarAutoupdate,
	9: ciscoboot,
}

var clo = map[int]uselessFunc{
	0: standard,
}

func main() {
    args := os.Args[1:]
	rand.Seed(time.Now().UnixNano())
	calls := rand.Intn(10)
	if len(args) > 0 {
		c, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Print("ok, you tried to pass the number of how (randomly) many program execute...\n")
			fmt.Printf("but what you type is not a number [%v]\n", args[0])
			fmt.Printf("so you will have the default for now :)\n")
		} else {
			calls = rand.Intn(c)
		}
	}
	fmt.Printf("here we go!  [%v]\n\n\n\n\n", calls)
	for i := 0; i < calls; i++ {
		call := vtp[rand.Intn(len(vtp))]
		call()
	}
	call := clo[rand.Intn(len(clo))]
	call()
}
