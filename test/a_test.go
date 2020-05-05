package main

import (
	"testing"
)

func Test_getHello(t *testing.T) {

	got := getHello()
	want := "hello world !!"

	if want != got {
		t.Errorf("failed test\nwant: %v\n got:%v", got, want)
		t.Fatal("failed test")
	}
}
