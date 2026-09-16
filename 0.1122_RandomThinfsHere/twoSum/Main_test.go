package twosum

import (
	"reflect"
	"testing"
)

func TestTwosomeBrutal(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Standard case - simple pair",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{2, 7}, // 2 + 7 = 9
		},
		{
			name:   "Target with different positions",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{2, 4}, // 2 + 4 = 6
		},
		{
			name:   "Allows using the same single index twice",
			nums:   []int{3, 2, 5},
			target: 6,
			want:   []int{3, 3}, // Because your code allows i == j, 3 + 3 = 6
		},
		{
			name:   "No solution exists",
			nums:   []int{1, 2, 3},
			target: 10,
			want:   nil,
		},
		{
			name:   "Empty slice",
			nums:   []int{},
			target: 5,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twosumBrutal(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twosomeBrutal() = %v, want %v", got, tt.want)
			}
		})
	}
}
