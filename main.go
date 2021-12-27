package main

import (
	"fmt"
	"math/rand"
	"os"
  	"math"
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
	9: pfSenseMenu,
	10: vm370,
	11: ciscoboot,
}

var clo = map[int]uselessFunc{
	0: standard,
}

func find(slice []int, val int) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func main() {
	args := os.Args[1:]
	rand.Seed(time.Now().UnixNano())
	calls := 0
	var err error
	if len(args) > 0 {
		calls, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Print("ok, you tried to pass the number of how many program execute...\n")
			fmt.Printf("but what you type is not a number [%v]\n", args[0])
			fmt.Printf("so you will have the default for now :)\n")
		}
  	}
  	calls = int(math.Max(float64(calls), float64(len(vtp))))
	if calls == 0 {
		calls = 1 + rand.Intn(len(vtp) - 1)
	}
	fmt.Printf("here we go!  [%v]\n\n\n\n\n", calls)
	
	// create execution list
  	myList := []int{}
	
	for i := 0; i < calls; {
		randomNumber := rand.Intn(len(vtp))
		_, found := find(myList, randomNumber)
		if !found {
			myList = append(myList, randomNumber)
			i++
		}
	}
	for i := 0; i < len(myList); i++ {
		call := vtp[myList[i]]
		call()
	}
	call := clo[rand.Intn(len(clo))]
	call()
}

