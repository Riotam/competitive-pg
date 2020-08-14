package main

import "testing"

func Test_checker(t *testing.T) {
	type args struct {
		count int
		line  []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"1",
			args{
				3,
				[]string{
					"SET 1 10",
					"SET 2 20",
					"ADD 40",
				},
			},
			10,
			50,
		},

		{
			"2",
			args{
				3,
				[]string{
					"SET 1 -23",
					"SUB 77",
					"SET 1 0",
				},
			},
			0,
			-100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checker(tt.args.count, tt.args.line)
			if got != tt.want {
				t.Errorf("checker() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checker() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
