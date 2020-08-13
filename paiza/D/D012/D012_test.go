package main

import "testing"

func Test_checker(t *testing.T) {
	type args struct {
		inputInt int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"value over than 0",
			args{inputInt: 100},
			100,
		},
		{
			"value less than 0",
			args{inputInt: -15},
			15,
		},
		{
			"value equal 0",
			args{inputInt: 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker(tt.args.inputInt); got != tt.want {
				t.Errorf("checker() = %v, want %v", got, tt.want)
			}
		})
	}
}
