package main

import "testing"

func Test_checker(t *testing.T) {

	input1 := []string{"chokudai", "chokudaiz"}
	got := checker(input1)
	want := "Yes"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input2 := []string{"snuke", "snekee"}
	got = checker(input2)
	want = "No"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	input3 := []string{"a", "aa"}
	got = checker(input3)
	want = "Yes"
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	// input4 := []string{"chokudai", "chokudaiz"}
	// got = checker(input4)
	// want = "NO"
	// if want != got {
	// 	t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	// }

	// input5 := []string{"chokudai", "chokudaiz"}
	// got = checker(input5)
	// want = "YES"
	// if want != got {
	// 	t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	// }

	// input6 := "rase"
	// got = checker(input6)
	// want = "NO"
	// if want != got {
	// 	t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	// }

}
