package main

import "testing"

func Test_checker(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"入力例1",
			args{
				"pa",
				"iza",
			},
			83,
		},
		{
			"入力例2",
			args{
				"alice",
				"bob",
			},
			97,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("checker() = %v, want %v", got, tt.want)
			}
		})
	}
}
