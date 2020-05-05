package util

import (
	"testing"
)

func Test_getNumbersByStr(t *testing.T) {
	s := "1 2 3"
	want := []int{1, 2, 3}
	got := getNumbersByStr(s)

	if len(got) == 0 {
		t.Errorf("failed test\nwant: %v\n got: %v", got, want)
	}

	for i, w := range want {
		if w != got[i] {
			t.Errorf("failed test\nwant: %v\n got: %v", got, want)
		}

	}
}
