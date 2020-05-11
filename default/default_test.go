package main

import "testing"

func Test_checker(t *testing.T) {

	input := "input"
	got := checker(input)
	want := "wantRes"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	// add test

}
