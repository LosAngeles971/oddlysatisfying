package main

import (
	"testing"
)

func TestMapOfFunc(t *testing.T) {
	for _, call := range vtp {
		t.Logf("testing logging function %v", call)
		call()
	}
}