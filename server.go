package main

import (
	"math/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type Loglines struct {
	Lines string
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./html/console.js.t")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	tt, err := template.New("loglines").Parse(string(b))
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	rand.Seed(time.Now().UnixNano())
    call := vtp[rand.Intn(len(vtp))]
	ll := call()
	ll.merge(standard())
	tt.Execute(w, ll)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	// mapping /console.js to the func in charge of generating it
	r.HandleFunc("/console.js", handler).Methods("GET")
	// where are static html resources
	staticFileDirectory := http.Dir("./html/")
	// how to handle static resources
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")
	return r
}