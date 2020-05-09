package main

import (
	"testing"
)

func Test_checker(t *testing.T) {

	input1 := "9 45000"
	got := checker(input1)
	want := "4 0 5"
	// 正解多数
	// if want != got {
	// 	t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	// }

	input2 := "20 196000"
	got = checker(input2)
	want = "-1 -1 -1"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input3 := "1000 1234000"
	got = checker(input3)
	want = "14 27 959"
	// 正解多数
	// if want != got {
	// 	t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	// }

	input4 := "2000 20000000"
	got = checker(input4)
	want = "2000 0 0"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

}
