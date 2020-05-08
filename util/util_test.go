package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_sumDigits(t *testing.T) {

	numbers := []int{11, 10, 19, 29, 100, 199, 1000, 19999}
	wants := []int{2, 1, 10, 11, 1, 19, 1, 37}

	for i, n := range numbers {
		got := sumDigits(n)

		if wants[i] != got {
			t.Errorf("failed test\nwant: %v\n got: %v", got, wants[i])
		}
	}

}

func Test_sortDesc(t *testing.T) {
	target := []int{1, 2, 3, 4, 0}
	sortDesc(target)
	want := []int{4, 3, 2, 1, 0}

	if !assert.Equal(t, target, want) {
		t.Errorf("failed test\ntarget: %v\n  want: %v", target, want)
	}
}
