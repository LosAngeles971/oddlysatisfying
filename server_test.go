package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDynamicFiles(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/console.js")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("wrong status code %d", resp.StatusCode)
	}
}

func TestStaticFiles(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/theme.css")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("wrong status code %d", resp.StatusCode)
	}
}