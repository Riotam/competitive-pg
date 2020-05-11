package main

import "testing"

func Test_checker(t *testing.T) {

	firstLineSliceInt := []int{3, 3, 10}
	lines := []string{"60 2 2 4", "70 8 7 9", "50 2 3 9"}
	got := checker(firstLineSliceInt, lines)
	want := 120
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

}
