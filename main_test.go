package main

import (
	"reflect"
	"testing"
)

func TestMapOfFuncIntegrity(t *testing.T) {
	for i, call := range vtp {
		tt := reflect.TypeOf(call).Kind()
		if tt != reflect.Func {
			t.Logf("element %d of map is not a function but %v", i, tt)
		}
	}
}