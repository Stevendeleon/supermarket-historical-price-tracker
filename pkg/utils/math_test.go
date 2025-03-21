package utils

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"empty input",
			args{[]int{}},
			0,
		},
		{
			"single number",
			args{[]int{1}},
			1,
		},
		{
			"multiple numbers",
			args{[]int{1, 2, 3}},
			6,
		},
		{
			"negative numbers",
			args{[]int{-1, -2, -3}},
			-6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.nums...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
