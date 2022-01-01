package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/jessevdk/go-flags"
)

const (
    server_http_listener = ":8080"
)

var opts struct {
	Server bool `short:"s" long:"server" description:"Start in server mode"`
    Port int `short:"p" long:"port" description:"Port used by server mode"`
    Logs int `short:"l" long:"logs" description:"How many logs for CLI mode"`
}

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

func toScreen(logs int) {
    rand.Seed(time.Now().UnixNano())
    calls := rand.Intn(len(vtp))
    if logs > 0 {
        calls = int(math.Min(float64(calls), float64(len(vtp))))
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
    _, err := flags.Parse(&opts)
    if err != nil {
        log.Fatal(err)
    }
    if opts.Server {
        r := newRouter()
        if opts.Port > 0 && opts.Port < 65536 {
            http.ListenAndServe(fmt.Sprintf(":%v", opts.Port), r)
        } else {
            http.ListenAndServe(server_http_listener, r)
        }
    } else {
        toScreen(opts.Logs)
    }
}
