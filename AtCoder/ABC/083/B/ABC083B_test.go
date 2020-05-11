package main

import "testing"

func Test_sumDigits(t *testing.T) {

	numbers := []int{11, 10, 19, 29, 100, 199}
	wants := []int{2, 1, 10, 11, 1, 19}

	for i, n := range numbers {
		got := sumDigits(n)

		if wants[i] != got {
			t.Errorf("failed test\nwant: %v\n got: %v", got, wants[i])
		}
	}

}
