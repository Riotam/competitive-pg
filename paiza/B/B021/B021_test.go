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
			"1",
			args{
				3,
				[]string{"dog", "cat", "pig"},
			},
			[]string{"dogs", "cats", "pigs"},
		},
		{
			"2",
			args{
				7,
				[]string{"box", "photo", "axis", "dish", "church", "leaf", "knife"},
			},
			[]string{"boxes", "photoes", "axises", "dishes", "churches", "leaves", "knives"},
		},
		{
			"3",
			args{
				2,
				[]string{"study", "play"},
			},
			[]string{"studies", "plays"},
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
