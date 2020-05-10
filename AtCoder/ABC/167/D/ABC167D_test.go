package main

import "testing"

func Test_checker(t *testing.T) {

	input1 := "erasedream"
	got := checker(input1)
	want := "YES"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input2 := "dreameraser"
	got = checker(input2)
	want = "YES"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input3 := "dreamerer"
	got = checker(input3)
	want = "NO"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input4 := "eraseraser"
	got = checker(input4)
	want = "NO"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input5 := "dreamererase"
	got = checker(input5)
	want = "YES"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input6 := "rase"
	got = checker(input6)
	want = "NO"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

}
