package main

import (
	"reflect"
	"testing"
)

func Test_checker(t *testing.T) {
	type args struct {
		count int
		line  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"入力例1",
			args{
				3,
				[]string{"28", "16", "777"},
			},
			[]string{
				"perfect",
				"nearly",
				"neither",
			},
		},
		{
			"入力例2",
			args{
				4,
				[]string{"3", "4", "5", "6"},
			},
			[]string{
				"neither",
				"nearly",
				"neither",
				"perfect",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checker(tt.args.count, tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checker() = %v, want %v", got, tt.want)
			}
		})
	}
}
