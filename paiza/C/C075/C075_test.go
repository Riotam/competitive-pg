package main

import (
	"reflect"
	"testing"
)

func Test_checker(t *testing.T) {
	type args struct {
		x    int
		y    int
		line []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"入力例1",
			args{
				2000,
				5,
				[]int{300, 500, 300, 100, 100},
			},
			[][]int{
				{1700, 30},
				{1200, 80},
				{900, 110},
				{900, 10},
				{800, 20},
			},
		},
		{
			"入力例2",
			args{
				3000,
				3,
				[]int{1000, 1000, 1000},
			},
			[][]int{
				{2000, 100},
				{1000, 200},
				{0, 300},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker(tt.args.x, tt.args.y, tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checker() = %v, want %v", got, tt.want)
			}
		})
	}
}
