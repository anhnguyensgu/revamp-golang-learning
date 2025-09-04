package p084_largest_rectangle_in_histogram

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		// Basic examples from original test
		{
			name:    "mixed_heights_1",
			heights: []int{1, 2, 3, 2, 2, 3, 3},
			want:    12, // width 4 * height 2 from indices 1-4
		},
		{
			name:    "leetcode_example",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10, // width 2 * height 5 from indices 2-3
		},
		{
			name:    "two_ascending",
			heights: []int{2, 4},
			want:    4, // single bar of height 4
		},

		// Edge cases
		{
			name:    "single_bar",
			heights: []int{5},
			want:    5,
		},
		{
			name:    "all_same_height",
			heights: []int{3, 3, 3, 3},
			want:    12, // width 4 * height 3
		},

		// Ascending sequences
		{
			name:    "strictly_ascending",
			heights: []int{1, 2, 3, 4, 5},
			want:    9, // width 3 * height 3 from middle
		},
		{
			name:    "two_bars_ascending",
			heights: []int{1, 3},
			want:    3,
		},

		// Descending sequences
		{
			name:    "strictly_descending",
			heights: []int{5, 4, 3, 2, 1},
			want:    9, // width 3 * height 3 from middle
		},
		{
			name:    "two_bars_descending",
			heights: []int{4, 2},
			want:    4,
		},

		// Special patterns
		{
			name:    "mountain_pattern",
			heights: []int{1, 3, 5, 3, 1},
			want:    9, // width 3 * height 3 from indices 1-3
		},
		{
			name:    "valley_pattern",
			heights: []int{5, 1, 5},
			want:    5, // single bar of height 5
		},
		{
			name:    "plateau_pattern",
			heights: []int{1, 4, 4, 4, 1},
			want:    12, // width 3 * height 4
		},

		// Zero heights
		{
			name:    "with_zero_heights",
			heights: []int{2, 0, 2},
			want:    2,
		},
		{
			name:    "all_zero_heights",
			heights: []int{0, 0, 0},
			want:    0,
		},

		// Complex patterns
		{
			name:    "multiple_peaks",
			heights: []int{2, 1, 5, 6, 2, 3, 1, 4},
			want:    10, // width 2 * height 5 from indices 2-3
		},
		{
			name:    "alternating_high_low",
			heights: []int{6, 1, 6, 1, 6},
			want:    6,
		},
		{
			name:    "large_rectangle_start",
			heights: []int{6, 6, 2, 1, 3},
			want:    12, // width 2 * height 6
		},
		{
			name:    "large_rectangle_end",
			heights: []int{2, 1, 3, 6, 6},
			want:    12, // width 2 * height 6
		},

		// Edge cases with large values
		{
			name:    "max_height_single",
			heights: []int{10000},
			want:    10000,
		},
		{
			name:    "increasing_then_decreasing",
			heights: []int{1, 2, 3, 4, 3, 2, 1},
			want:    16, // width 4 * height 4 or other optimal
		},

		// Stress test patterns
		{
			name:    "long_uniform",
			heights: []int{5, 5, 5, 5, 5, 5, 5, 5},
			want:    40, // width 8 * height 5
		},
		{
			name:    "zigzag_pattern",
			heights: []int{1, 3, 1, 3, 1, 3, 1},
			want:    7,
		},
		{
			name:    "zigzag_pattern",
			heights: []int{2, 1, 2},
			want:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.heights)
			if got != tt.want {
				t.Errorf("largestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
			}
		})
	}
}
