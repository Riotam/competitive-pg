package main

import "testing"

func Test_checker(t *testing.T) {
	type args struct {
		inputList [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"入力例1",
			args{
				[][]int{
					{30, 30, 10},
					{3},
					{6, 0},
					{7, 0},
					{8, 0},
				},
			},
			"07:30",
		},
		{
			"入力例2",
			args{
				[][]int{
					{30, 30, 10},
					{3},
					{6, 0},
					{7, 0},
					{8, 0},
				},
			},
			"07:30",
		},
		{
			"入力例3",
			args{
				[][]int{
					{25, 30, 30},
					{2},
					{7, 20},
					{8, 0},
				},
			},
			"06:55",
		},
		{
			"入力例4",
			args{
				[][]int{
					{10, 10, 10},
					{6},
					{8, 5},
					{8, 15},
					{8, 25},
					{8, 35},
					{8, 45},
					{8, 55},
				},
			},
			"08:25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker(tt.args.inputList); got != tt.want {
				t.Errorf("checker() = %v, want %v", got, tt.want)
			}
		})
	}
}
