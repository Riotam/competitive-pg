package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_example1(t *testing.T) {

	// 1
	int1 := 6
	word1 := "a"
	got1 := example(int1, word1)
	want1 := "over five"

	if !assert.Equal(t, got1, want1) {
		t.Errorf("failed test\nword: %v\n   got: %v", word1, got1)
	}

	// 2
	int2 := 5
	word2 := "a"
	got2 := example(int2, word2)
	want2 := "over zero"

	if !assert.Equal(t, got2, want2) {
		t.Errorf("failed test\nword: %v\n   got: %v", word2, got2)
	}

	// 3
	int3 := 1
	word3 := "a"
	got3 := example(int3, word3)
	want3 := "over zero"

	if !assert.Equal(t, got3, want3) {
		t.Errorf("failed test\nword: %v\n   got: %v", word3, got3)
	}

	// 4
	int4 := 0
	word4 := "a"
	got4 := example(int4, word4)
	want4 := "too small"

	if !assert.Equal(t, got4, want4) {
		t.Errorf("failed test\nword: %v\n   got: %v", word4, got4)
	}

	// 5
	int5 := 0
	word5 := "a"
	got5 := example(int5, word5)
	want5 := "too small"

	if !assert.Equal(t, got5, want5) {
		t.Errorf("failed test\nword: %v\n   got: %v", word5, got5)
	}

	// 6
	int6 := -1
	word6 := ""
	got6 := example(int6, word6)
	want6 := "wrong word"

	if !assert.Equal(t, got6, want6) {
		t.Errorf("failed test\nword: %v\n   got: %v", word6, got6)
	}
}
