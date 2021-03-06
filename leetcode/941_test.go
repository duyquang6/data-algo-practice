package leetcode

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				A: []int{2, 1},
			},
			want: false,
		},
		{
			args: args{
				A: []int{3, 5, 5},
			},
			want: false,
		},
		{
			args: args{
				A: []int{3, 5, 6},
			},
			want: false,
		},
		{
			args: args{
				A: []int{0, 3, 2, 1},
			},
			want: true,
		},
		{
			args: args{
				A: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
