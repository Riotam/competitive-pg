package main

import "testing"

func Test_checker(t *testing.T) {

	a := 2
	b := 1
	c := 1
	k := 3
	want := 2
	got := checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a2 := 1
	b2 := 2
	c2 := 3
	k2 := 4
	want2 := 0
	got2 := checker(a2, b2, c2, k2)
	if want2 != got2 {
		t.Errorf("failed test\n got: %v\nwant: %v", got2, want2)
	}

	a3 := 2000000000
	b3 := 0
	c3 := 0
	k3 := 2000000000
	want3 := 2000000000
	got3 := checker(a3, b3, c3, k3)
	if want3 != got3 {
		t.Errorf("failed test\n got: %v\nwant: %v", got3, want3)
	}

	a4 := 3
	b4 := 0
	c4 := 3
	k4 := 4
	want4 := 2
	got4 := checker(a4, b4, c4, k4)
	if want4 != got4 {
		t.Errorf("failed test\n got: %v\nwant: %v", got4, want4)
	}

	a = 100
	b = 0
	c = 0
	k = 100
	want = 100
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a = 1
	b = 0
	c = 100
	k = 100
	want = -98
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a = 1
	b = 2
	c = 1
	k = 3
	want = 1
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a = 0
	b = 0
	c = 0
	k = 0
	want = 0
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a = 2
	b = 2
	c = 2
	k = 5
	want = 1
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a = 0
	b = 3
	c = 3
	k = 4
	want = -1
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a = 10
	b = 0
	c = 0
	k = 5
	want = 5
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

	a = 0
	b = 0
	c = 10
	k = 5
	want = -5
	got = checker(a, b, c, k)
	if want != got {
		t.Errorf("failed test\n got: %v\nwant: %v", got, want)
	}

}
