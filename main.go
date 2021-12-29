package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
    server_http_listener = ":8080"
)

type uselessFunc func() *Log

var vtp = map[int]uselessFunc{
    0:  download,
    1:  mvs,
    2:  bootVax,
    3:  c64,
    4:  adminports,
    5:  fsck,
    6:  bootMsDos,
    7:  stateLinux,
    8:  qradarAutoupdate,
    9:  pfSenseMenu,
    10: vm370,
    11: ciscoboot,
    12: curlDownload,
}

var clo = map[int]uselessFunc{
    0: standard,
}

func toScreen(args []string) {
    rand.Seed(time.Now().UnixNano())
    calls := rand.Intn(len(vtp))
    var err error
    if len(args) > 0 {
        calls = 0
        calls, err = strconv.Atoi(args[0])
        if err != nil {
            fmt.Print("ok, you tried to pass the number of how many program execute...\n")
            fmt.Printf("but what you type is not a number [%v]\n", args[0])
            fmt.Printf("so you will have the default for now :)\n")
        }
    }
    calls = int(math.Min(float64(calls), float64(len(vtp))))
    if calls == 0 {
        calls = 1
    }
    fmt.Printf("here we go!  [%v]\n\n\n\n\n", calls)
    myList := []int{}
    for i := 0; i < len(vtp); i++ {
        myList = append(myList, i)
    }
    rand.Shuffle(len(myList), func(i, j int) { myList[i], myList[j] = myList[j], myList[i] })
    myList = myList[:calls]
    for i := 0; i < len(myList); i++ {
        call := vtp[myList[i]]
        log := call()
        log.toScreen()
    }
    call := clo[rand.Intn(len(clo))]
    call()
}

func main() {
    args := os.Args[1:]
    if len(args) > 0 && args[0] == "server" {
        r := newRouter()
        http.ListenAndServe(server_http_listener, r)
    } else {
        toScreen(args)
    }
}
