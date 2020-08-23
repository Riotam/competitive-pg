package main

import (
	"reflect"
	"testing"
)

func Test_checker(t *testing.T) {
	type args struct {
		lineList [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"入力例1",
			args{
				[][]int{
					{7, 10, 4},
					{1, 8, 1},
					{4, 1, 5},
					{1, 6, 2},
					{2, 2, 0},
				},
			},
			[]string{
				"..........",
				"..######..",
				".....#....",
				".....#....",
				"##...#....",
				"##...#....",
				".########.",
			},
		},
		{
			"入力例2",
			args{
				[][]int{
					{10, 10, 9},
					{2, 2, 4},
					{2, 2, 3},
					{2, 2, 5},
					{2, 2, 2},
					{2, 2, 6},
					{2, 2, 1},
					{2, 2, 7},
					{2, 2, 0},
					{2, 2, 8},
				},
			},
			[]string{
				"##......##",
				"##......##",
				".##....##.",
				".##....##.",
				"..##..##..",
				"..##..##..",
				"...####...",
				"...####...",
				"....##....",
				"....##....",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker(tt.args.lineList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checker() = %v, want %v", got, tt.want)
			}
		})
	}
}
