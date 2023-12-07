package main

import "testing"

func Test_intersect(t *testing.T) {
	type args struct {
		range1 []int
		range2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains full intersection",
			args: args{
				range1: []int{0, 4},
				range2: []int{0, 4},
			},
			want: true,
		},
		{
			name: "contains subset",
			args: args{
				range1: []int{0, 4},
				range2: []int{1, 3},
			},
			want: true,
		},
		{
			name: "Not contains left boundary",
			args: args{
				range1: []int{0, 3},
				range2: []int{3, 6},
			},
			want: false,
		},
		{
			name: "contains left boundary intersection",
			args: args{
				range1: []int{0, 4},
				range2: []int{3, 6},
			},
			want: true,
		},
		{
			name: "Not contains right boundary",
			args: args{
				range1: []int{6, 8},
				range2: []int{3, 6},
			},
			want: false,
		},
		{
			name: "Not contains right boundary exactly",
			args: args{
				range1: []int{6, 10},
				range2: []int{3, 6},
			},
			want: false,
		},
		{
			name: "contains right boundary intersection",
			args: args{
				range1: []int{5, 10},
				range2: []int{3, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.range1, tt.args.range2); got != tt.want {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
