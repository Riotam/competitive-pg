package main

import "testing"

func Test_checkEven(t *testing.T) {

	a := 3
	b := 4
	want := "Even"
	got := checkEven(a, b)

	if want != got {
		t.Errorf("failed test\nwant: %v\n got: %v", got, want)
	}

	a = 1
	b = 21
	want = "Odd"
	got = checkEven(a, b)

	if want != got {
		t.Errorf("failed test\nwant: %v\n got: %v", got, want)
	}

	a = 3
	b = 5
	want = "Odd"
	got = checkEven(a, b)

	if want != got {
		t.Errorf("failed test\nwant: %v\n got: %v", got, want)
	}

}
